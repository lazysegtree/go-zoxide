package zoxide

import "fmt"

func (c *Client) Add(path string) error {
	zArgs := getZargsForAdd(path, false, "")
	_, err := c.execCmd(zArgs...)

	// TODO: Improve error handling
	if err != nil {
		return fmt.Errorf("add failed due to unexpected command execution error: %w", err)
	}
	return nil
}

func getZargsForAdd(path string, score bool, scoreValue string) []string {
	zArgs := []string{"add"}
	if score {
		zArgs = append(zArgs, "-s", scoreValue)
	}
	zArgs = append(zArgs, path)
	return zArgs
}
