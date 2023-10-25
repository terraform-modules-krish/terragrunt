package outputmodulegroups

import (
	"fmt"

	"github.com/terraform-modules-krish/terragrunt/cli/commands/terraform"
	"github.com/terraform-modules-krish/terragrunt/config"
	"github.com/terraform-modules-krish/terragrunt/configstack"
	"github.com/terraform-modules-krish/terragrunt/options"
)

func Run(opts *options.TerragruntOptions) error {
	target := terraform.NewTarget(terraform.TargetPointParseConfig, runOutputModuleGroups)

	return terraform.RunWithTarget(opts, target)
}

func runOutputModuleGroups(opts *options.TerragruntOptions, cfg *config.TerragruntConfig) error {
	stack, err := configstack.FindStackInSubfolders(opts, nil)
	if err != nil {
		return err
	}

	js, err := stack.JsonModuleDeployOrder(opts.TerraformCommand)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(opts.Writer, "%s\n", js)
	if err != nil {
		return err
	}

	return nil

}
