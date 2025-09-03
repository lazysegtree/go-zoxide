package zoxide

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
)

func (c *Client) Query(args ...string) (string, error) {
	res, err := c.QueryWithOptions(false, false, false, args...)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		// Use defined error type
		return "", errors.New("no results found")
	}
	return res[0].Path, nil
}

func (c *Client) QueryAll(args ...string) ([]Result, error) {
	return c.QueryWithOptions(true, true, true, args...)
}

func (c *Client) QueryWithOptions(all bool, list bool, score bool,
	args ...string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	zArgs := append(getZargsForQuery(all, list, score), args...)
	cmd := exec.CommandContext(ctx, commandName, zArgs...)

	stdout, err := cmd.Output()

	// TODO: Handle timeout error and no match found error
	// And return them correctly in Query()
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return nil, fmt.Errorf("query command exited with code %d : %w\nstderr : %v",
				exitError.ExitCode(), err, string(exitError.Stderr))
		}
		return nil, fmt.Errorf("query failed due to unexpected command execution error of : %w", err)
	}

	return parseResults(string(stdout), score)
}

func getZargsForQuery(all bool, list bool, score bool) []string {
	zargs := []string{"query"}
	if all {
		zargs = append(zargs, "-a")
	}
	if list {
		zargs = append(zargs, "-l")
	}
	if score {
		zargs = append(zargs, "-s")
	}
	return zargs
}
