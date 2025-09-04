// Package zoxide provides a Go wrapper for the zoxide command-line tool,
// enabling integration with zoxide's smarter directory navigation in Go applications.
//
// zoxide is a smarter cd command that learns your habits and helps you jump
// to frequently and recently used directories. This package provides a clean
// API to interact with zoxide from Go programs.
//
// Basic usage:
//
//	client, err := zoxide.New()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Query for the best match
//	path, err := client.Query("myproject")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("Best match:", path)
//
//	// Get all matches with scores
//	results, err := client.QueryAll("project")
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, result := range results {
//		fmt.Printf("%.2f %s\n", result.Score, result.Path)
//	}
package zoxide

import (
	"os/exec"
	"time"
)

// Client represents a zoxide client with configurable options.
// It manages communication with the zoxide command-line tool and handles
// execution timeouts and custom data directories.
type Client struct {
	timeout time.Duration // execution timeout for zoxide commands
	dataDir string        // custom path to zoxide data directory (optional)
}

// ClientOption is a function type used to configure Client instances.
// Options can be passed to New() to customize client behavior.
type ClientOption func(*Client) error

// New creates a new zoxide client with the given options.
// The client is initialized with default settings and can be customized
// using ClientOption functions.
//
// Example:
//
//	// Create client with default settings
//	client, err := zoxide.New()
//
//	// Create client with custom data directory
//	client, err := zoxide.New(zoxide.WithDataDir("/custom/path"))
//
// Returns an error if any of the provided options fail to apply.
func New(opts ...ClientOption) (*Client, error) {
	if _, err := exec.LookPath(commandName); err != nil {
		return nil, ErrZoxideNotFound
	}
	c := &Client{
		timeout: defaultExecTimeout,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithDataDir returns a ClientOption that sets a custom data directory for zoxide.
// This allows you to use a different zoxide database than the default system-wide one.
//
// The path should be an absolute path to a directory where zoxide will store its data.
// If the directory doesn't exist, zoxide will create it when needed.
//
// Example:
//
//	client, err := zoxide.New(zoxide.WithDataDir("/home/user/.local/share/custom-zoxide"))
func WithDataDir(path string) ClientOption {
	return func(c *Client) error {
		// TODO : validate path is absolute
		c.dataDir = path
		return nil
	}
}
