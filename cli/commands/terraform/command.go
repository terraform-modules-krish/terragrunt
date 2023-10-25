package terraform

import (
	"github.com/terraform-modules-krish/terragrunt/options"
	"github.com/terraform-modules-krish/terragrunt/pkg/cli"
)

const (
	CommandName = "terraform"
)

func NewCommand(opts *options.TerragruntOptions) *cli.Command {
	return &cli.Command{
		Name:     CommandName,
		HelpName: "*",
		Usage:    "Terragrunt forwards all other commands directly to Terraform",
		Action:   action(opts),
	}
}

func action(opts *options.TerragruntOptions) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) error {
		if opts.TerraformCommand == CommandNameDestroy {
			opts.CheckDependentModules = true
		}

		return Run(opts.OptionsFromContext(ctx))
	}
}
