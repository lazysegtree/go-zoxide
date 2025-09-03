package zoxide

import (
	"time"
)

type Client struct {
	timeout time.Duration
	dataDir string
}

type ClientOption func(*Client) error

func New(opts ...ClientOption) (*Client, error) {
	// TODO: Check if zoxide exists in PATH, or return nil
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

func WithDataDir(path string) ClientOption {
	return func(c *Client) error {
		// TODO : validate path is absolute
		c.dataDir = path
		return nil
	}
}
