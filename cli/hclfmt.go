package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/mattn/go-zglob"

	"github.com/terraform-modules-krish/terragrunt/errors"
	"github.com/terraform-modules-krish/terragrunt/options"
	"github.com/terraform-modules-krish/terragrunt/util"
)

// runHCLFmt recursively looks for hcl files in the directory tree starting at workingDir, and formats them
// based on the language style guides provided by Hashicorp. This is done using the official hcl2 library.
func runHCLFmt(terragruntOptions *options.TerragruntOptions) error {

	workingDir := terragruntOptions.WorkingDir
	targetFile := terragruntOptions.HclFile

	// handle when option specifies a particular file
	if targetFile != "" {
		if !filepath.IsAbs(targetFile) {
			targetFile = util.JoinPath(workingDir, targetFile)
		}
		terragruntOptions.Logger.Infof("Formatting hcl file at: %s.", targetFile)
		return formatTgHCL(terragruntOptions, targetFile)
	}

	terragruntOptions.Logger.Infof("Formatting hcl files from the directory tree %s.", terragruntOptions.WorkingDir)
	// zglob normalizes paths to "/"
	tgHclFiles, err := zglob.Glob(util.JoinPath(workingDir, "**", "*.hcl"))
	if err != nil {
		return err
	}

	filteredTgHclFiles := []string{}
	for _, fname := range tgHclFiles {
		// Ignore any files that are in the .terragrunt-cache
		if !util.ListContainsElement(strings.Split(fname, "/"), ".terragrunt-cache") {
			filteredTgHclFiles = append(filteredTgHclFiles, fname)
		} else {
			terragruntOptions.Logger.Debugf("%s was ignored due to being in the terragrunt cache", fname)
		}
	}

	terragruntOptions.Logger.Debugf("Found %d hcl files", len(filteredTgHclFiles))

	formatErrors := []error{}
	for _, tgHclFile := range filteredTgHclFiles {
		err := formatTgHCL(terragruntOptions, tgHclFile)
		if err != nil {
			formatErrors = append(formatErrors, err)
		}
	}

	return errors.NewMultiError(formatErrors...)
}

// formatTgHCL uses the hcl2 library to format the hcl file. This will attempt to parse the HCL file first to
// ensure that there are no syntax errors, before attempting to format it.
func formatTgHCL(terragruntOptions *options.TerragruntOptions, tgHclFile string) error {
	terragruntOptions.Logger.Infof("Formatting %s", tgHclFile)

	info, err := os.Stat(tgHclFile)
	if err != nil {
		terragruntOptions.Logger.Errorf("Error retrieving file info of %s", tgHclFile)
		return err
	}

	contentsStr, err := util.ReadFileAsString(tgHclFile)
	if err != nil {
		terragruntOptions.Logger.Errorf("Error reading %s", tgHclFile)
		return err
	}
	contents := []byte(contentsStr)

	err = checkErrors(contents, tgHclFile)
	if err != nil {
		terragruntOptions.Logger.Errorf("Error parsing %s", tgHclFile)
		return err
	}

	newContents := hclwrite.Format(contents)

	if terragruntOptions.Check {
		if !bytes.Equal(newContents, contents) {
			return fmt.Errorf("Invalid file format %s", tgHclFile)
		}
		return nil
	}

	return ioutil.WriteFile(tgHclFile, newContents, info.Mode())
}

// checkErrors takes in the contents of a hcl file and looks for syntax errors.
func checkErrors(contents []byte, tgHclFile string) error {
	parser := hclparse.NewParser()
	_, diags := parser.ParseHCL(contents, tgHclFile)
	diagWriter := util.GetDiagnosticsWriter(parser)
	diagWriter.WriteDiagnostics(diags)
	if diags.HasErrors() {
		return diags
	}
	return nil
}
