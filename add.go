package zoxide

import "fmt"

// Add adds a directory path to the zoxide database.
// This increments the directory's score, making it more likely to appear
// in future queries. The path should be an absolute path to a directory.
//
// This is equivalent to running `zoxide add <path>`.
//
// Example:
//
//	err := client.Add("/home/user/projects/myproject")
//	if err != nil {
//		log.Fatal(err)
//	}
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
