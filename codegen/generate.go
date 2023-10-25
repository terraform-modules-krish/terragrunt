package codegen

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
	ctyjson "github.com/zclconf/go-cty/cty/json"

	"github.com/terraform-modules-krish/go-commons/errors"
	"github.com/terraform-modules-krish/terragrunt/options"
	"github.com/terraform-modules-krish/terragrunt/util"
)

const (
	// A comment that is added to the top of the generated file to indicate that this file was generated by Terragrunt.
	// We use a hardcoded random string at the end to make the string further unique.
	TerragruntGeneratedSignature = "Generated by Terragrunt. Sig: nIlQXj57tbuaRZEa"

	// The default prefix to use for comments in the generated file
	DefaultCommentPrefix = "# "
)

// An enum to represent valid values for if_exists
type GenerateConfigExists int

const (
	ExistsError GenerateConfigExists = iota
	ExistsSkip
	ExistsOverwrite
	ExistsOverwriteTerragrunt
	ExistsUnknown
)

const (
	ExistsErrorStr               = "error"
	ExistsSkipStr                = "skip"
	ExistsOverwriteStr           = "overwrite"
	ExistsOverwriteTerragruntStr = "overwrite_terragrunt"
)

// Configuration for generating code
type GenerateConfig struct {
	Path             string `cty:"path"`
	IfExists         GenerateConfigExists
	IfExistsStr      string `cty:"if_exists"`
	CommentPrefix    string `cty:"comment_prefix"`
	Contents         string `cty:"contents"`
	DisableSignature bool   `cty:"disable_signature"`
	Disable          bool   `cty:"disable"`
}

// WriteToFile will generate a new file at the given target path with the given contents. If a file already exists at
// the target path, the behavior depends on the value of IfExists:
// - if ExistsError, return an error.
// - if ExistsSkip, do nothing and return
// - if ExistsOverwrite, overwrite the existing file
func WriteToFile(terragruntOptions *options.TerragruntOptions, basePath string, config GenerateConfig) error {
	// If this GenerateConfig is disabled then skip further processing.
	if config.Disable {
		terragruntOptions.Logger.Debugf("Skipping generating file at %s because it is disabled", config.Path)
		return nil
	}

	// Figure out thee target path to generate the code in. If relative, merge with basePath.
	var targetPath string
	if filepath.IsAbs(config.Path) {
		targetPath = config.Path
	} else {
		targetPath = filepath.Join(basePath, config.Path)
	}

	targetFileExists := util.FileExists(targetPath)
	if targetFileExists {
		shouldContinue, err := shouldContinueWithFileExists(terragruntOptions, targetPath, config.IfExists)
		if err != nil || !shouldContinue {
			return err
		}
	}

	// Add the signature as a prefix to the file, unless it is disabled.
	prefix := ""
	if !config.DisableSignature {
		prefix = fmt.Sprintf("%s%s\n", config.CommentPrefix, TerragruntGeneratedSignature)
	}
	contentsToWrite := fmt.Sprintf("%s%s", prefix, config.Contents)

	if err := ioutil.WriteFile(targetPath, []byte(contentsToWrite), 0644); err != nil {
		return errors.WithStackTrace(err)
	}
	terragruntOptions.Logger.Debugf("Generated file %s.", targetPath)
	return nil
}

// Whether or not file generation should continue if the file path already exists. The answer depends on the
// ifExists configuration.
func shouldContinueWithFileExists(terragruntOptions *options.TerragruntOptions, path string, ifExists GenerateConfigExists) (bool, error) {
	switch ifExists {
	case ExistsError:
		return false, errors.WithStackTrace(GenerateFileExistsError{path: path})
	case ExistsSkip:
		// Do nothing since file exists and skip was configured
		terragruntOptions.Logger.Debugf("The file path %s already exists and if_exists for code generation set to \"skip\". Will not regenerate file.", path)
		return false, nil
	case ExistsOverwrite:
		// We will continue to proceed to generate file, but log a message to indicate that we detected the file
		// exists.
		terragruntOptions.Logger.Debugf("The file path %s already exists and if_exists for code generation set to \"overwrite\". Regenerating file.", path)
		return true, nil
	case ExistsOverwriteTerragrunt:
		// If file was not generated, error out because overwrite_terragrunt if_exists setting only handles if the
		// existing file was generated by terragrunt.
		wasGenerated, err := fileWasGeneratedByTerragrunt(path)
		if err != nil {
			return false, err
		}
		if !wasGenerated {
			terragruntOptions.Logger.Errorf("ERROR: The file path %s already exists and was not generated by terragrunt.", path)
			return false, errors.WithStackTrace(GenerateFileExistsError{path: path})
		}
		// Since file was generated by terragrunt, continue.
		terragruntOptions.Logger.Debugf("The file path %s already exists, but was a previously generated file by terragrunt. Since if_exists for code generation is set to \"overwrite_terragrunt\", regenerating file.", path)
		return true, nil
	default:
		// This shouldn't happen, but we add this case anyway for defensive coding.
		return false, errors.WithStackTrace(UnknownGenerateIfExistsVal{""})
	}
}

// Check if the file was generated by terragrunt by checking if the first line of the file has the signature. Since the
// generated string will be prefixed with the configured comment prefix, the check needs to see if the first line ends
// with the signature string.
func fileWasGeneratedByTerragrunt(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, errors.WithStackTrace(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	firstLine, err := reader.ReadString('\n')
	if err != nil {
		return false, errors.WithStackTrace(err)
	}
	return strings.HasSuffix(strings.TrimSpace(firstLine), TerragruntGeneratedSignature), nil
}

// Convert the arbitrary map that represents a remote state config into HCL code to configure that remote state.
func RemoteStateConfigToTerraformCode(backend string, config map[string]interface{}) ([]byte, error) {
	f := hclwrite.NewEmptyFile()
	backendBlock := f.Body().AppendNewBlock("terraform", nil).Body().AppendNewBlock("backend", []string{backend})
	backendBlockBody := backendBlock.Body()
	var backendKeys []string

	for key := range config {
		backendKeys = append(backendKeys, key)
	}
	sort.Strings(backendKeys)
	for _, key := range backendKeys {
		// Since we don't have the cty type information for the config and since config can be arbitrary, we cheat by using
		// json as an intermediate representation.
		jsonBytes, err := json.Marshal(config[key])
		if err != nil {
			return nil, errors.WithStackTrace(err)
		}
		var ctyVal ctyjson.SimpleJSONValue
		if err := ctyVal.UnmarshalJSON(jsonBytes); err != nil {
			return nil, errors.WithStackTrace(err)
		}

		backendBlockBody.SetAttributeValue(key, ctyVal.Value)
	}

	return f.Bytes(), nil
}

// GenerateConfigExistsFromString converts a string representation of if_exists into the enum, returning an error if it
// is not set to one of the known values.
func GenerateConfigExistsFromString(val string) (GenerateConfigExists, error) {
	switch val {
	case ExistsErrorStr:
		return ExistsError, nil
	case ExistsSkipStr:
		return ExistsSkip, nil
	case ExistsOverwriteStr:
		return ExistsOverwrite, nil
	case ExistsOverwriteTerragruntStr:
		return ExistsOverwriteTerragrunt, nil
	}
	return ExistsUnknown, errors.WithStackTrace(UnknownGenerateIfExistsVal{val: val})
}

// Custom error types

type UnknownGenerateIfExistsVal struct {
	val string
}

func (err UnknownGenerateIfExistsVal) Error() string {
	if err.val != "" {
		return fmt.Sprintf("%s is not a valid value for generate if_exists", err.val)
	}
	return "Received unknown value for if_exists"
}

type GenerateFileExistsError struct {
	path string
}

func (err GenerateFileExistsError) Error() string {
	return fmt.Sprintf("Can not generate terraform file: %s already exists", err.path)
}
