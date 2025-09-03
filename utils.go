package zoxide

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func (c *Client) execCmd(zArgs ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, commandName, zArgs...)
	envVars := c.getEnvVars()
	if envVars != nil {
		cmd.Env = append(os.Environ(), envVars...)
	}
	stdout, err := cmd.Output()

	// TODO: Handle timeout error and no match found error
	// And return them correctly in Query()
	if err != nil {
		// TODO: Define a command failed error with extra data
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return nil, fmt.Errorf("command exited with code %d : %w\nstderr : %v",
				exitError.ExitCode(), err, string(exitError.Stderr))
		}
		return nil, fmt.Errorf("command failed due to unexpected error of : %w", err)
	}

	return stdout, nil
}

func (c *Client) getEnvVars() []string {
	var res []string
	if c.dataDir != "" {
		res = append(res, "_ZO_DATA_DIR="+c.dataDir)
	}

	return res
}
