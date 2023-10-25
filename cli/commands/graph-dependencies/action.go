package graphdependencies

import (
	"github.com/terraform-modules-krish/terragrunt/configstack"
	"github.com/terraform-modules-krish/terragrunt/options"
)

// Run graph dependencies prints the dependency graph to stdout
func Run(opts *options.TerragruntOptions) error {
	stack, err := configstack.FindStackInSubfolders(opts, nil)
	if err != nil {
		return err
	}

	// Exit early if the operation wanted is to get the graph
	stack.Graph(opts)
	return nil
}
