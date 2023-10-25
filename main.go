package main

import (
	"os"
	"github.com/terraform-modules-krish/terragrunt/util"
	"github.com/terraform-modules-krish/terragrunt/cli"
	"github.com/terraform-modules-krish/terragrunt/errors"
)

// This variable is set at build time using -ldflags parameters. For more info, see:
// http://stackoverflow.com/a/11355611/483528
var VERSION string

// The main entrypoint for Terragrunt
func main() {
	defer errors.Recover(checkForErrorsAndExit)

	app := cli.CreateTerragruntCli(VERSION)
	err := app.Run(os.Args)

	checkForErrorsAndExit(err)
}

// If there is an error, display it in the console and exit with a non-zero exit code. Otherwise, exit 0.
func checkForErrorsAndExit(err error) {
	if err == nil {
		os.Exit(0)
	} else {
		if os.Getenv("TERRAGRUNT_DEBUG") != "" {
			util.Logger.Println(errors.PrintErrorWithStackTrace(err))
		} else {
			util.Logger.Println(err)
		}
		os.Exit(1)
	}

}