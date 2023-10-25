package spin

import (
	"fmt"
	"strings"

	"github.com/terraform-modules-krish/terragrunt/config"
	"github.com/terraform-modules-krish/terragrunt/errors"
	"github.com/terraform-modules-krish/terragrunt/options"

	"github.com/mattn/go-zglob"
)

// Represents a stack of Terraform modules (i.e. folders with Terraform templates) that you can "spin up" or
// "spin down" in a single command
type Stack struct {
	Path    string
	Modules []*TerraformModule
}

// Render this stack as a human-readable string
func (stack *Stack) String() string {
	modules := []string{}
	for _, module := range stack.Modules {
		modules = append(modules, fmt.Sprintf("  => %s", module.String()))
	}
	return fmt.Sprintf("Stack at %s:\n%s", stack.Path, strings.Join(modules, "\n"))
}

// Apply all the modules in the given stack, making sure to apply the dependencies of each module in the stack in the
// proper order.
func (stack *Stack) Apply(terragruntOptions *options.TerragruntOptions) error {
	stack.setTerraformCommand([]string{"apply", "-input=false"})
	return RunModules(stack.Modules)
}

// Destroy all the modules in the given stack, making sure to destroy the dependencies of each module in the stack in
// the proper order.
func (stack *Stack) Destroy(terragruntOptions *options.TerragruntOptions) error {
	stack.setTerraformCommand([]string{"destroy", "-force", "-input=false"})
	return RunModulesReverseOrder(stack.Modules)
}

// Return an error if there is a dependency cycle in the modules of this stack.
func (stack *Stack) CheckForCycles() error {
	return CheckForCycles(stack.Modules)
}

// Find all the Terraform modules in the subfolders of the working directory of the given TerragruntOptions and
// assemble them into a Stack object that can be applied or destroyed in a single command
func FindStackInSubfolders(terragruntOptions *options.TerragruntOptions) (*Stack, error) {
	// Ideally, we'd use a builin Go library like filepath.Glob here, but per https://github.com/golang/go/issues/11862,
	// the current go implementation doesn't support treating ** as zero or more directories, just zero or one.
	// So we use a third-party library.
	terragruntConfigFiles, err := zglob.Glob(fmt.Sprintf("%s/**/%s", terragruntOptions.WorkingDir, config.DefaultTerragruntConfigPath))
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	return createStackForTerragruntConfigPaths(terragruntOptions.WorkingDir, terragruntConfigFiles, terragruntOptions)
}

// Set the command in the TerragruntOptions object of each module in this stack to the given command.
func (stack *Stack) setTerraformCommand(command []string) {
	for _, module := range stack.Modules {
		module.TerragruntOptions.TerraformCliArgs = append(command, module.TerragruntOptions.TerraformCliArgs...)
	}
}

// Find all the Terraform modules in the folders that contain the given Terragrunt config files and assemble those
// modules into a Stack object that can be applied or destroyed in a single command
func createStackForTerragruntConfigPaths(path string, terragruntConfigPaths []string, terragruntOptions *options.TerragruntOptions) (*Stack, error) {
	if len(terragruntConfigPaths) == 0 {
		return nil, errors.WithStackTrace(NoTerraformModulesFound)
	}

	modules, err := ResolveTerraformModules(terragruntConfigPaths, terragruntOptions)
	if err != nil {
		return nil, err
	}

	stack := &Stack{Path: path, Modules: modules}
	if err := stack.CheckForCycles(); err != nil {
		return nil, err
	}

	return stack, nil
}

// Custom error types

var NoTerraformModulesFound = fmt.Errorf("Could not find any subfolders with %s files", config.DefaultTerragruntConfigPath)

type DependencyCycle []string
func (err DependencyCycle) Error() string {
	return fmt.Sprintf("Found a dependency cycle between modules: %s", strings.Join([]string(err), " -> "))
}