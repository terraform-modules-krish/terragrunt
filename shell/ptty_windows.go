// +build windows

package shell

import (
	"io"
	"os"
	"os/exec"

	"github.com/terraform-modules-krish/terragrunt/errors"
	"github.com/terraform-modules-krish/terragrunt/options"
)

// For windows, there is no concept of a pseudoTTY so we run as if there is no pseudoTTY.
func runCommandWithPTTY(terragruntOptions *options.TerragruntOptions, cmd *exec.Cmd, cmdStdout io.Writer, cmdStderr io.Writer) error {
	cmd.Stdin = os.Stdin
	cmd.Stdout = cmdStdout
	cmd.Stderr = cmdStderr
	if err := cmd.Start(); err != nil {
		// bad path, binary not executable, &c
		return errors.WithStackTrace(err)
	}
	return nil
}
