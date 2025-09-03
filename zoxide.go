package zoxide

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	timeout time.Duration
}

func New() *Client {
	return &Client{
		timeout: defaultExecTimeout,
	}
}

type Result struct {
	Score float64
	Path  string
}

func (c *Client) Query(args ...string) ([]Result, error) {
	return c.QueryWithOptions(false, true, true, args...)
}

func (c *Client) QueryWithOptions(all bool, list bool, score bool,
	args ...string) ([]Result, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, commandName, getZargsForQuery(all, list, score)...)

	stdout, err := cmd.Output()

	// TODO: Handle timeout error and no match found error
	if err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			return nil, fmt.Errorf("query command exited with code %d : %w\nstderr : %v",
				exitError.ExitCode(), err, string(exitError.Stderr))
		}
		return nil, fmt.Errorf("query failed due to unexpected command execution error of : %w", err)
	}

	return parseResult(string(stdout), score)
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

func parseResult(stdout string, withScore bool) ([]Result, error) {
	res := []Result{}

	// TODO: Is this the most efficient way of doing this ?
	for line := range strings.Lines(stdout) {
		line = strings.TrimSpace(line)
		var score float64
		var path string
		if withScore {
			score, path = splitScoreAndPath(line)
		} else {
			score, path = defaultScore, line
		}
		res = append(res, Result{Score: score, Path: path})
	}
	return res, nil
}

func splitScoreAndPath(line string) (float64, string) {
	tokens := strings.SplitN(line, scoreSeperator, 2)
	// TODO: Length should be 2

	// TODO : Handle error
	score, _ := strconv.ParseFloat(tokens[0], 64)

	return score, strings.TrimSpace(tokens[1])
}
