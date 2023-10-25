package test

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/terraform-modules-krish/terragrunt/cli"
	"bytes"
	"time"
	"math/rand"
	"io/ioutil"
	"path"
	"github.com/terraform-modules-krish/terragrunt/remote"
	"github.com/stretchr/testify/assert"
)

// hard-code this to match the test fixture for now
const (
	TERRAFORM_REMOTE_STATE_S3_REGION      = "us-west-2"
	TEST_FIXTURE_PATH                     = "fixture/"
	TEST_FIXTURE_LOCK_PATH                = "fixture-lock/"
)

func TestTerragruntWorksWithLocalTerraformVersion(t *testing.T) {
	t.Parallel()

	s3BucketName := fmt.Sprintf("terragrunt-test-bucket-%s", strings.ToLower(uniqueId()))
	tmpTerragruntConfigPath := createTmpTerragruntConfig(t, TEST_FIXTURE_PATH, s3BucketName)

	defer deleteS3Bucket(t, TERRAFORM_REMOTE_STATE_S3_REGION, s3BucketName)
	runTerragrunt(t, fmt.Sprintf("terragrunty apply --terragrunt-non-interactive --terragrunt-config %s %s", tmpTerragruntConfigPath, TEST_FIXTURE_PATH))
	validateS3BucketExists(t, TERRAFORM_REMOTE_STATE_S3_REGION, s3BucketName)
}

func TestAcquireAndReleaseLock(t *testing.T) {
	t.Parallel()

	terragruntConfigPath := path.Join(TEST_FIXTURE_LOCK_PATH, ".terragrunt")

	// Acquire a long-term lock
	runTerragrunt(t, fmt.Sprintf("terragrunt acquire-lock --terragrunt-non-interactive --terragrunt-config %s", terragruntConfigPath))

	// Try to apply the templates. Since a lock has been acquired, and max_lock_retries is set to 1, this should
	// fail quickly.
	err := runTerragruntCommand(t, fmt.Sprintf("terragrunt apply --terragrunt-non-interactive --terragrunt-config %s %s", terragruntConfigPath, TEST_FIXTURE_LOCK_PATH))

	if assert.NotNil(t, err, "Expected to get an error when trying to apply templates after a long-term lock has already been acquired, but got nil") {
		assert.Contains(t, err.Error(), "Unable to acquire lock")
	}

	// Release the lock
	runTerragrunt(t, fmt.Sprintf("terragrunt release-lock --terragrunt-non-interactive --terragrunt-config %s", terragruntConfigPath))

	// Try to apply the templates. Since the lock has been released, this should work without errors.
	runTerragrunt(t, fmt.Sprintf("terragrunt apply --terragrunt-non-interactive --terragrunt-config %s %s", terragruntConfigPath, TEST_FIXTURE_LOCK_PATH))
}

func runTerragruntCommand(t *testing.T, command string) error {
	validateCommandInstalled(t, "terraform")
	args := strings.Split(command, " ")

	app := cli.CreateTerragruntCli("TEST")
	return app.Run(args)
}

func runTerragrunt(t *testing.T, command string) {
	if err := runTerragruntCommand(t, command); err != nil {
		t.Fatalf("Failed to run Terragrunt command '%s' due to error: %s", command, err)
	}
}

func createTmpTerragruntConfig(t *testing.T, templatesPath string, s3BucketName string) string {
	tmpTerragruntConfigFile, err := ioutil.TempFile("", ".terragrunt")
	if err != nil {
		t.Fatalf("Failed to create temp file due to error: %v", err)
	}

	originalTerragruntConfigPath := path.Join(templatesPath, ".terragrunt")
	originalTerragruntConfigBytes, err := ioutil.ReadFile(originalTerragruntConfigPath)
	if err != nil {
		t.Fatalf("Error reading Terragrunt config at %s: %v", originalTerragruntConfigPath, err)
	}

	originalTerragruntConfigString := string(originalTerragruntConfigBytes)
	newTerragruntConfigString := strings.Replace(originalTerragruntConfigString, "__FILL_IN_BUCKET_NAME__", s3BucketName, -1)

	if err := ioutil.WriteFile(tmpTerragruntConfigFile.Name(), []byte(newTerragruntConfigString), 0444); err != nil {
		t.Fatalf("Error writing temp Terragrunt config to %s: %v", tmpTerragruntConfigFile.Name(), err)
	}

	return tmpTerragruntConfigFile.Name()
}

// Returns a unique (ish) id we can attach to resources and tfstate files so they don't conflict with each other
// Uses base 62 to generate a 6 character string that's unlikely to collide with the handful of tests we run in
// parallel. Based on code here: http://stackoverflow.com/a/9543797/483528
func uniqueId() string {
	const BASE_62_CHARS = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	const UNIQUE_ID_LENGTH = 6 // Should be good for 62^6 = 56+ billion combinations

	var out bytes.Buffer

	randInstance := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < UNIQUE_ID_LENGTH; i++ {
		out.WriteByte(BASE_62_CHARS[randInstance.Intn(len(BASE_62_CHARS))])
	}

	return out.String()
}

// Validate that the given command is available in PATH
func validateCommandInstalled(t *testing.T, command string) {
	_, err := exec.LookPath(command)
	if err != nil {
		t.Fatalf("Command '%s' not found in PATH", command)
	}
}

// Check that the S3 Bucket of the given name and region exists. Terragrunt should create this bucket during the test.
func validateS3BucketExists(t *testing.T, awsRegion string, bucketName string) {
	s3Client, err := remote.CreateS3Client(awsRegion)
	if err != nil {
		t.Fatalf("Error creating S3 client: %v", err)
	}

	remoteStateConfig := remote.RemoteStateConfigS3{Bucket: bucketName, Region: awsRegion}
	assert.True(t, remote.DoesS3BucketExist(s3Client, &remoteStateConfig), "Terragrunt failed to create remote state S3 bucket %s", bucketName)
}

// Delete the specified S3 bucket to clean up after a test
func deleteS3Bucket(t *testing.T, awsRegion string, bucketName string) {
	s3Client, err := remote.CreateS3Client(awsRegion)
	if err != nil {
		t.Fatalf("Error creating S3 client: %v", err)
	}

	t.Logf("Deleting test s3 bucket %s", bucketName)

	out, err := s3Client.ListObjectVersions(&s3.ListObjectVersionsInput{Bucket: aws.String(bucketName)})
	if err != nil {
		t.Fatalf("Failed to list object versions in s3 bucket %s: %v", bucketName, err)
	}

	objectIdentifiers := []*s3.ObjectIdentifier{}
	for _, version := range out.Versions {
		objectIdentifiers = append(objectIdentifiers, &s3.ObjectIdentifier{
			Key: version.Key,
			VersionId: version.VersionId,
		})
	}

	deleteInput := &s3.DeleteObjectsInput{
		Bucket: aws.String(bucketName),
		Delete: &s3.Delete{Objects: objectIdentifiers},
	}
	if _, err := s3Client.DeleteObjects(deleteInput); err != nil {
		t.Fatalf("Error deleting all versions of all objects in bucket %s: %v", bucketName, err)
	}

	if _, err := s3Client.DeleteBucket(&s3.DeleteBucketInput{Bucket: aws.String(bucketName)}); err != nil {
		t.Fatalf("Failed to delete S3 bucket %s: %v", bucketName, err)
	}
}
