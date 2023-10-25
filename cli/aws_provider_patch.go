package cli

import (
	"encoding/json"
	"fmt"
	"github.com/terraform-modules-krish/terragrunt/errors"
	"github.com/terraform-modules-krish/terragrunt/options"
	"github.com/terraform-modules-krish/terragrunt/util"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/mattn/go-zglob"
	"github.com/zclconf/go-cty/cty"
	"io/ioutil"
	"path/filepath"
)

// applyAwsProviderPatch finds all Terraform modules nested in the current code (i.e., in the .terraform/modules
// folder), looks for provider "aws" { ... } blocks in those modules, and overwrites the attributes in those provider
// blocks with the attributes specified in terragrntOptions.
//
// For example, if were running Terragrunt against code that contained a module:
//
// module "example" {
//   source = "<URL>"
// }
//
// When you run 'init', Terraform would download the code for that module into .terraform/modules. This function would
// scan that module code for provider blocks:
//
// provider "aws" {
//    region = var.aws_region
// }
//
// And if AwsProviderPatchOverrides in terragruntOptions was set to map[string]string{"region": "us-east-1"}, then this
// method would update the module code to:
//
// provider "aws" {
//    region = "us-east-1"
// }
//
// This is a temporary workaround for a Terraform bug (https://github.com/hashicorp/terraform/issues/13018) where
// any dynamic values in nested provider blocks are not handled correctly when you call 'terraform import', so by
// temporarily hard-coding them, we can allow 'import' to work.
func applyAwsProviderPatch(terragruntOptions *options.TerragruntOptions) error {
	if len(terragruntOptions.AwsProviderPatchOverrides) == 0 {
		return errors.WithStackTrace(MissingOverrides(OPT_TERRAGRUNT_OVERRIDE_ATTR))
	}

	terraformFilesInModules, err := findAllTerraformFilesInModules(terragruntOptions)
	if err != nil {
		return err
	}

	for _, terraformFile := range terraformFilesInModules {
		util.Debugf(terragruntOptions.Logger, "Looking at file %s", terraformFile)
		originalTerraformFileContents, err := util.ReadFileAsString(terraformFile)
		if err != nil {
			return err
		}

		updatedTerraformFileContents, codeWasUpdated, err := patchAwsProviderInTerraformCode(originalTerraformFileContents, terraformFile, terragruntOptions.AwsProviderPatchOverrides)
		if err != nil {
			return err
		}

		if codeWasUpdated {
			terragruntOptions.Logger.Printf("Patching AWS provider in %s", terraformFile)
			if err := util.WriteFileWithSamePermissions(terraformFile, terraformFile, []byte(updatedTerraformFileContents)); err != nil {
				return err
			}
		}
	}

	return nil
}

// The format we expect in the .terraform/modules/modules.json file
type TerraformModulesJson struct {
	Modules []TerraformModule
}

type TerraformModule struct {
	Key    string
	Source string
	Dir    string
}

// findAllTerraformFiles returns all Terraform source files within the modules being used by this Terragrunt
// configuration. To be more specific, it only returns the source files downloaded for module "xxx" { ... } blocks into
// the .terraform/modules folder; it does NOT return Terraform files for the top-level (AKA "root") module.
//
// NOTE: this method only supports *.tf files right now. Terraform code defined in *.json files is not currently
// supported.
func findAllTerraformFilesInModules(terragruntOptions *options.TerragruntOptions) ([]string, error) {
	// Terraform downloads modules into the .terraform/modules folder. Unfortunately, it downloads not only the module
	// into that folder, but the entire repo it's in, which can contain lots of other unrelated code we probably don't
	// want to touch. To find the paths to the actual modules, we read the modules.json file in that folder, which is
	// a manifest file Terraform uses to track where the modules are within each repo. Note that this is an internal
	// API, so the way we parse/read this modules.json file may break in future Terraform versions. Note that we
	// can't use the official HashiCorp code to parse this file, as it's marked internal:
	// https://github.com/hashicorp/terraform/blob/master/internal/modsdir/manifest.go
	modulesJsonPath := util.JoinPath(terragruntOptions.DataDir(), "modules", "modules.json")

	if !util.FileExists(modulesJsonPath) {
		return nil, nil
	}

	modulesJsonContents, err := ioutil.ReadFile(modulesJsonPath)
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	var terraformModulesJson TerraformModulesJson
	if err := json.Unmarshal(modulesJsonContents, &terraformModulesJson); err != nil {
		return nil, errors.WithStackTrace(err)
	}

	var terraformFiles []string

	for _, module := range terraformModulesJson.Modules {
		if module.Key != "" && module.Dir != "" {
			moduleAbsPath := module.Dir
			if !filepath.IsAbs(moduleAbsPath) {
				moduleAbsPath = util.JoinPath(terragruntOptions.WorkingDir, moduleAbsPath)
			}

			// Ideally, we'd use a builtin Go library like filepath.Glob here, but per https://github.com/golang/go/issues/11862,
			// the current go implementation doesn't support treating ** as zero or more directories, just zero or one.
			// So we use a third-party library.
			matches, err := zglob.Glob(fmt.Sprintf("%s/**/*.tf", moduleAbsPath))
			if err != nil {
				return nil, errors.WithStackTrace(err)
			}

			terraformFiles = append(terraformFiles, matches...)
		}
	}

	return terraformFiles, nil
}

// patchAwsProviderInTerraformCode looks for provider "aws" { ... } blocks in the given Terraform code and overwrites
// the attributes in those provider blocks with the given attributes. It returns the new Terraform code and a boolean
// true if that code was updated.
//
// For example, if you passed in the following Terraform code:
//
// provider "aws" {
//    region = var.aws_region
// }
//
// And you set attributesToOverride to map[string]string{"region": "us-east-1"}, then this method will return:
//
// provider "aws" {
//    region = "us-east-1"
// }
//
// This is a temporary workaround for a Terraform bug (https://github.com/hashicorp/terraform/issues/13018) where
// any dynamic values in nested provider blocks are not handled correctly when you call 'terraform import', so by
// temporarily hard-coding them, we can allow 'import' to work.
func patchAwsProviderInTerraformCode(terraformCode string, terraformFilePath string, attributesToOverride map[string]string) (string, bool, error) {
	if len(attributesToOverride) == 0 {
		return terraformCode, false, nil
	}

	hclFile, err := hclwrite.ParseConfig([]byte(terraformCode), terraformFilePath, hcl.InitialPos)
	if err != nil {
		return "", false, errors.WithStackTrace(err)
	}

	codeWasUpdated := false

	for _, block := range hclFile.Body().Blocks() {
		if block.Type() == "provider" && len(block.Labels()) == 1 && block.Labels()[0] == "aws" {
			for key, value := range attributesToOverride {
				block.Body().SetAttributeValue(key, cty.StringVal(value))
			}

			codeWasUpdated = true
		}
	}

	return string(hclFile.Bytes()), codeWasUpdated, nil
}

// Custom error types

type MissingOverrides string

func (err MissingOverrides) Error() string {
	return fmt.Sprintf("You must specify at least one provider attribute to override via the --%s option.", string(err))
}
