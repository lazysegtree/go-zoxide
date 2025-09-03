package zoxide

import (
	"errors"
	"fmt"
)

// Query searches zoxide for the best matching directory based on the given arguments.
// It returns the single best match as a string path.
//
// This is equivalent to running `zoxide query <args>` and getting the top result.
// If no matches are found, returns an error.
//
// Example:
//
//	path, err := client.Query("myproject")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Best match:", path)
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

// QueryAll searches zoxide and returns all matching directories with their scores.
// It returns a slice of Result structs containing both paths and scores.
//
// This is equivalent to running `zoxide query -a -l -s <args>` and parsing
// the results. Results are ordered by score (highest first).
//
// Example:
//
//	results, err := client.QueryAll("project")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, result := range results {
//		fmt.Printf("%.2f %s\n", result.Score, result.Path)
//	}
func (c *Client) QueryAll(args ...string) ([]Result, error) {
	return c.QueryWithOptions(true, true, true, args...)
}

// QueryWithOptions provides fine-grained control over zoxide query behavior.
// This is a lower-level method that allows customizing query flags.
//
// Parameters:
//   - all: If true, show all results (equivalent to -a flag)
//   - list: If true, list results (equivalent to -l flag) 
//   - score: If true, include scores in results (equivalent to -s flag)
//   - args: Search terms to pass to zoxide
//
// Most users should use Query() or QueryAll() instead of this method.
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
