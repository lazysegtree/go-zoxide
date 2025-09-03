package zoxide

import (
	"time"
)

type Client struct {
	timeout time.Duration
}

func New() (*Client, error) {
	// TODO: Check if zoxide exists in PATH, or return nil
	return &Client{
		timeout: defaultExecTimeout,
	}, nil
}
