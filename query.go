package zoxide

import (
	"errors"
	"fmt"
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

	zArgs := append(getZargsForQuery(all, list, score), args...)

	stdout, err := c.execCmd(zArgs...)

	// TODO: Improve error handling
	if err != nil {
		return nil, fmt.Errorf("query failed due to unexpected command execution error : %w", err)
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
