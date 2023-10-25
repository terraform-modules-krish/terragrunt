package config

import (
	"reflect"
	"testing"
	"github.com/terraform-modules-krish/terragrunt/errors"
	"github.com/terraform-modules-krish/terragrunt/locks/dynamodb"
	"github.com/terraform-modules-krish/terragrunt/remote"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-modules-krish/terragrunt/options"
	"github.com/terraform-modules-krish/terragrunt/util"
)

var mockOptions = options.TerragruntOptions{TerragruntConfigPath: "test-time-mock", NonInteractive: true}

func TestParseTerragruntConfigDynamoLockMinimalConfig(t *testing.T) {
	t.Parallel()

	config :=
`
lock = {
  backend = "dynamodb"
  config {
    state_file_id = "expected-state-file-id"
  }
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Terraform)
	assert.NotNil(t, terragruntConfig.Lock)
	assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
	lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
	assert.Equal(t, "expected-state-file-id", lock.StateFileId)
	assert.Equal(t, dynamodb.DEFAULT_AWS_REGION, lock.AwsRegion)
	assert.Equal(t, dynamodb.DEFAULT_TABLE_NAME, lock.TableName)
	assert.Equal(t, dynamodb.DEFAULT_MAX_RETRIES_WAITING_FOR_LOCK, lock.MaxLockRetries)
}

func TestParseTerragruntConfigDynamoLockFullConfig(t *testing.T) {
	t.Parallel()

	config :=
`
lock = {
  backend = "dynamodb"
  config {
    state_file_id = "expected-state-file-id"
    aws_region = "expected-region"
    table_name = "expected-table-name"
    max_lock_retries = 100
  }
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Terraform)

	if assert.NotNil(t, terragruntConfig.Lock) {
		assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
		lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
		assert.Equal(t, "expected-state-file-id", lock.StateFileId)
		assert.Equal(t, "expected-region", lock.AwsRegion)
		assert.Equal(t, "expected-table-name", lock.TableName)
		assert.Equal(t, 100, lock.MaxLockRetries)
	}
}

func TestParseTerragruntConfigDynamoLockMissingStateFileId(t *testing.T) {
	t.Parallel()

	config :=
`
lock = {
  backend = "dynamodb"
  config {
  }
}
`

	_, err := parseConfigString(config, &mockOptions, nil)
	assert.True(t, errors.IsError(err, dynamodb.StateFileIdMissing))
}

func TestParseTerragruntConfigRemoteStateMinimalConfig(t *testing.T) {
	t.Parallel()

	config :=
`
remote_state = {
  backend = "s3"
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.Lock)
	assert.Nil(t, terragruntConfig.Terraform)

	if assert.NotNil(t, terragruntConfig.RemoteState){
		assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
		assert.Empty(t, terragruntConfig.RemoteState.Config)
	}
}

func TestParseTerragruntConfigRemoteStateMissingBackend(t *testing.T) {
	t.Parallel()

	config :=
`
remote_state = {
}
`

	_, err := parseConfigString(config, &mockOptions, nil)
	assert.True(t, errors.IsError(err, remote.RemoteBackendMissing), "Unexpected error of type %s: %s", reflect.TypeOf(err), err)
}

func TestParseTerragruntConfigRemoteStateFullConfig(t *testing.T) {
	t.Parallel()

	config :=
`
remote_state = {
  backend = "s3"
  config = {
    encrypt = "true"
    bucket = "my-bucket"
    key = "terraform.tfstate"
    region = "us-east-1"
  }
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.Lock)
	assert.Nil(t, terragruntConfig.Terraform)

	if assert.NotNil(t, terragruntConfig.RemoteState) {
		assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
		assert.NotEmpty(t, terragruntConfig.RemoteState.Config)
		assert.Equal(t, "true", terragruntConfig.RemoteState.Config["encrypt"])
		assert.Equal(t, "my-bucket", terragruntConfig.RemoteState.Config["bucket"])
		assert.Equal(t, "terraform.tfstate", terragruntConfig.RemoteState.Config["key"])
		assert.Equal(t, "us-east-1", terragruntConfig.RemoteState.Config["region"])
	}
}

func TestParseTerragruntConfigDependenciesOnePath(t *testing.T) {
	t.Parallel()

	config :=
`
dependencies = {
  paths = ["../vpc"]
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.Lock)
	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Terraform)

	if assert.NotNil(t, terragruntConfig.Dependencies) {
		assert.Equal(t, []string{"../vpc"}, terragruntConfig.Dependencies.Paths)
	}
}

func TestParseTerragruntConfigDependenciesMultiplePaths(t *testing.T) {
	t.Parallel()

	config :=
`
dependencies = {
  paths = ["../vpc", "../mysql", "../backend-app"]
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.Lock)
	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Terraform)

	if assert.NotNil(t, terragruntConfig.Dependencies) {
		assert.Equal(t, []string{"../vpc", "../mysql", "../backend-app"}, terragruntConfig.Dependencies.Paths)
	}
}

func TestParseTerragruntConfigRemoteStateDynamoDbTerraformConfigAndDependenciesFullConfig(t *testing.T) {
	t.Parallel()

	config :=
`
terraform {
  source = "foo"
}

lock {
  backend = "dynamodb"
  config {
    state_file_id = "expected-state-file-id"
    aws_region = "expected-region"
    table_name = "expected-table-name"
    max_lock_retries = 100
  }
}

remote_state {
  backend = "s3"
  config {
    encrypt = "true"
    bucket = "my-bucket"
    key = "terraform.tfstate"
    region = "us-east-1"
  }
}

dependencies = {
  paths = ["../vpc", "../mysql", "../backend-app"]
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	if assert.NotNil(t, terragruntConfig.Terraform) {
		assert.Equal(t, "foo", terragruntConfig.Terraform.Source)
	}

	if assert.NotNil(t, terragruntConfig.Lock) {
		assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
		lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
		assert.Equal(t, "expected-state-file-id", lock.StateFileId)
		assert.Equal(t, "expected-region", lock.AwsRegion)
		assert.Equal(t, "expected-table-name", lock.TableName)
		assert.Equal(t, 100, lock.MaxLockRetries)
	}

	if assert.NotNil(t, terragruntConfig.RemoteState) {
		assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
		assert.NotEmpty(t, terragruntConfig.RemoteState.Config)
		assert.Equal(t, "true", terragruntConfig.RemoteState.Config["encrypt"])
		assert.Equal(t, "my-bucket", terragruntConfig.RemoteState.Config["bucket"])
		assert.Equal(t, "terraform.tfstate", terragruntConfig.RemoteState.Config["key"])
		assert.Equal(t, "us-east-1", terragruntConfig.RemoteState.Config["region"])
	}

	if assert.NotNil(t, terragruntConfig.Dependencies) {
		assert.Equal(t, []string{"../vpc", "../mysql", "../backend-app"}, terragruntConfig.Dependencies.Paths)
	}
}

func TestParseTerragruntConfigInvalidLockBackend(t *testing.T) {
	t.Parallel()

	config :=
`
lock = {
  backend = "invalid"
  config {
  }
}
`

	_, err := parseConfigString(config, &mockOptions, nil)
	assert.True(t, errors.IsError(err, ErrLockNotFound))
}

func TestParseTerragruntConfigInclude(t *testing.T) {
	t.Parallel()

	config :=
`
include = {
  path = "../../../.terragrunt"
}
`

	opts := options.TerragruntOptions{
		TerragruntConfigPath: "../test/fixture-parent-folders/terragrunt-in-root/child/sub-child/sub-sub-child/.terragrunt",
		NonInteractive: true,
	}

	terragruntConfig, err := parseConfigString(config, &opts, nil)
	if assert.Nil(t, err, "Unexpected error: %v", errors.PrintErrorWithStackTrace(err)) {
		assert.Nil(t, terragruntConfig.Terraform)

		if assert.NotNil(t, terragruntConfig.Lock) {
			assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
			lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
			assert.Equal(t, "child/sub-child/sub-sub-child", lock.StateFileId)
		}

		if assert.NotNil(t, terragruntConfig.RemoteState) {
			assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
			assert.NotEmpty(t, terragruntConfig.RemoteState.Config)
			assert.Equal(t, "true", terragruntConfig.RemoteState.Config["encrypt"])
			assert.Equal(t, "my-bucket", terragruntConfig.RemoteState.Config["bucket"])
			assert.Equal(t, "child/sub-child/sub-sub-child/terraform.tfstate", terragruntConfig.RemoteState.Config["key"])
			assert.Equal(t, "us-east-1", terragruntConfig.RemoteState.Config["region"])
		}
	}

}

func TestParseTerragruntConfigIncludeWithFindInParentFolders(t *testing.T) {
	t.Parallel()

	config :=
`
include = {
  path = "${find_in_parent_folders()}"
}
`

	opts := options.TerragruntOptions{
		TerragruntConfigPath: "../test/fixture-parent-folders/terragrunt-in-root/child/sub-child/sub-sub-child/.terragrunt",
		NonInteractive: true,
	}

	terragruntConfig, err := parseConfigString(config, &opts, nil)
	if assert.Nil(t, err, "Unexpected error: %v", errors.PrintErrorWithStackTrace(err)) {
		assert.Nil(t, terragruntConfig.Terraform)

		if assert.NotNil(t, terragruntConfig.Lock) {
			assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
			lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
			assert.Equal(t, "child/sub-child/sub-sub-child", lock.StateFileId)
		}


		if assert.NotNil(t, terragruntConfig.RemoteState) {
			assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
			assert.NotEmpty(t, terragruntConfig.RemoteState.Config)
			assert.Equal(t, "true", terragruntConfig.RemoteState.Config["encrypt"])
			assert.Equal(t, "my-bucket", terragruntConfig.RemoteState.Config["bucket"])
			assert.Equal(t, "child/sub-child/sub-sub-child/terraform.tfstate", terragruntConfig.RemoteState.Config["key"])
			assert.Equal(t, "us-east-1", terragruntConfig.RemoteState.Config["region"])
		}
	}

}

func TestParseTerragruntConfigIncludeOverrideRemote(t *testing.T) {
	t.Parallel()

	config :=
`
include = {
  path = "../../../.terragrunt"
}

# Configure Terragrunt to automatically store tfstate files in an S3 bucket
remote_state = {
  backend = "s3"
  config {
    encrypt = "override"
    bucket = "override"
    key = "override"
    region = "override"
  }
}
`

	opts := options.TerragruntOptions{
		TerragruntConfigPath: "../test/fixture-parent-folders/terragrunt-in-root/child/sub-child/sub-sub-child/.terragrunt",
		NonInteractive: true,
	}

	terragruntConfig, err := parseConfigString(config, &opts, nil)
	if assert.Nil(t, err, "Unexpected error: %v", errors.PrintErrorWithStackTrace(err)) {
		assert.Nil(t, terragruntConfig.Terraform)

		if assert.NotNil(t, terragruntConfig.Lock) {
			assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
			lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
			assert.Equal(t, "child/sub-child/sub-sub-child", lock.StateFileId)
		}

		if assert.NotNil(t, terragruntConfig.RemoteState) {
			assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
			assert.NotEmpty(t, terragruntConfig.RemoteState.Config)
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["encrypt"])
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["bucket"])
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["key"])
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["region"])
		}
	}

}

func TestParseTerragruntConfigIncludeOverrideAll(t *testing.T) {
	t.Parallel()

	config :=
`
include {
  path = "../../../.terragrunt"
}

terraform {
  source = "foo"
}

lock = {
  backend = "dynamodb"
  config {
    state_file_id = "override"
  }
}

# Configure Terragrunt to automatically store tfstate files in an S3 bucket
remote_state {
  backend = "s3"
  config {
    encrypt = "override"
    bucket = "override"
    key = "override"
    region = "override"
  }
}

dependencies {
  paths = ["override"]
}
`

	opts := options.TerragruntOptions{
		TerragruntConfigPath: "../test/fixture-parent-folders/terragrunt-in-root/child/sub-child/sub-sub-child/.terragrunt",
		NonInteractive: true,
	}

	terragruntConfig, err := parseConfigString(config, &opts, nil)
	if assert.Nil(t, err, "Unexpected error: %v", errors.PrintErrorWithStackTrace(err)) {
		if assert.NotNil(t, terragruntConfig.Terraform) {
			assert.Equal(t, "foo", terragruntConfig.Terraform.Source)
		}

		if assert.NotNil(t, terragruntConfig.Lock) {
			assert.IsType(t, &dynamodb.DynamoDbLock{}, terragruntConfig.Lock)
			lock := terragruntConfig.Lock.(*dynamodb.DynamoDbLock)
			assert.Equal(t, "override", lock.StateFileId)
		}

		if assert.NotNil(t, terragruntConfig.RemoteState) {
			assert.Equal(t, "s3", terragruntConfig.RemoteState.Backend)
			assert.NotEmpty(t, terragruntConfig.RemoteState.Config)
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["encrypt"])
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["bucket"])
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["key"])
			assert.Equal(t, "override", terragruntConfig.RemoteState.Config["region"])
		}

		assert.Equal(t, []string{"override"}, terragruntConfig.Dependencies.Paths)
	}

}

func TestParseTerragruntConfigTwoLevels(t *testing.T) {
	t.Parallel()

	configPath := "../test/fixture-parent-folders/multiple-terragrunt-in-parents/child/sub-child/.terragrunt"

	config, err := util.ReadFileAsString(configPath)
	if err != nil {
		t.Fatal(err)
	}

	opts := options.TerragruntOptions{TerragruntConfigPath: configPath, NonInteractive: true}

	_, actualErr := parseConfigString(config, &opts, nil)
	expectedErr := TooManyLevelsOfInheritance(configPath)
	assert.True(t, errors.IsError(actualErr, expectedErr), "Expected error %v but got %v", expectedErr, actualErr)
}

func TestParseTerragruntConfigThreeLevels(t *testing.T) {
	t.Parallel()

	configPath := "../test/fixture-parent-folders/multiple-terragrunt-in-parents/child/sub-child/sub-sub-child/.terragrunt"

	config, err := util.ReadFileAsString(configPath)
	if err != nil {
		t.Fatal(err)
	}

	opts := options.TerragruntOptions{TerragruntConfigPath: configPath, NonInteractive: true}

	_, actualErr := parseConfigString(config, &opts, nil)
	expectedErr := TooManyLevelsOfInheritance(configPath)
	assert.True(t, errors.IsError(actualErr, expectedErr), "Expected error %v but got %v", expectedErr, actualErr)
}

func TestParseTerragruntConfigEmptyConfig(t *testing.T) {
	t.Parallel()

	config := ``

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Lock)
}

func TestMergeConfigIntoIncludedConfig(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		config         *TerragruntConfig
		includedConfig *TerragruntConfig
		expected       *TerragruntConfig
	}{
		{
			&TerragruntConfig{},
			nil,
			&TerragruntConfig{},
		},
		{
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}},
			nil,
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}},
		},
		{
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}, RemoteState: &remote.RemoteState{Backend: "foo"}},
			nil,
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}, RemoteState: &remote.RemoteState{Backend: "foo"}},
		},
		{
			&TerragruntConfig{},
			&TerragruntConfig{},
			&TerragruntConfig{},
		},
		{
			&TerragruntConfig{},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}},
		},
		{
			&TerragruntConfig{},
			&TerragruntConfig{Terraform: &TerraformConfig{Source: "foo"}},
			&TerragruntConfig{Terraform: &TerraformConfig{Source: "foo"}},
		},
		{
			&TerragruntConfig{},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "foo"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "foo"}},
		},
		{
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "foo"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "foo"}},
		},
		{
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}, RemoteState: &remote.RemoteState{Backend: "foo"}, Terraform: &TerraformConfig{Source: "foo"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "bar"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "foo"}, RemoteState: &remote.RemoteState{Backend: "foo"}, Terraform: &TerraformConfig{Source: "foo"}},
		},
		{
			&TerragruntConfig{RemoteState: &remote.RemoteState{Backend: "foo"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "bar"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "foo"}, Terraform: &TerraformConfig{Source: "bar"}},
		},
		{
			&TerragruntConfig{Terraform: &TerraformConfig{Source: "foo"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "bar"}},
			&TerragruntConfig{Lock: dynamodb.DynamoDbLock{StateFileId: "bar"}, RemoteState: &remote.RemoteState{Backend: "bar"}, Terraform: &TerraformConfig{Source: "foo"}},
		},
	}

	for _, testCase := range testCases {
		actual, err := mergeConfigWithIncludedConfig(testCase.config, testCase.includedConfig)
		if assert.Nil(t, err, "Unexpected error for config %v and includeConfig %v: %v", testCase.config, testCase.includedConfig, err) {
			assert.Equal(t, testCase.expected, actual, "For config %v and includeConfig %v", testCase.config, testCase.includedConfig)
		}
	}
}

func TestParseTerragruntConfigTerraformNoSource(t *testing.T) {
	t.Parallel()

	config :=
`
terraform {
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.Lock)
	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Dependencies)

	if assert.NotNil(t, terragruntConfig.Terraform) {
		assert.Empty(t, terragruntConfig.Terraform.Source)
	}
}

func TestParseTerragruntConfigTerraformWithSource(t *testing.T) {
	t.Parallel()

	config :=
`
terraform {
  source = "foo"
}
`

	terragruntConfig, err := parseConfigString(config, &mockOptions, nil)
	assert.Nil(t, err)

	assert.Nil(t, terragruntConfig.Lock)
	assert.Nil(t, terragruntConfig.RemoteState)
	assert.Nil(t, terragruntConfig.Dependencies)

	if assert.NotNil(t, terragruntConfig.Terraform) {
		assert.Equal(t, "foo", terragruntConfig.Terraform.Source)
	}
}