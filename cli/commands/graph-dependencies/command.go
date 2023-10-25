package graphdependencies

import (
	"github.com/terraform-modules-krish/terragrunt/options"
	"github.com/terraform-modules-krish/terragrunt/pkg/cli"
)

const (
	CommandName = "graph-dependencies"
)

func NewCommand(opts *options.TerragruntOptions) *cli.Command {
	return &cli.Command{
		Name:   CommandName,
		Usage:  "Prints the terragrunt dependency graph to stdout.",
		Action: func(ctx *cli.Context) error { return Run(opts.OptionsFromContext(ctx)) },
	}
}
