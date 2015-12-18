package awsuserdata

import (
	"strings"
)

// Client provides a shell for the env client
type Client struct{}

// NewAwsClient returns a new client
func NewAwsUserDataClient() (*Client, error) {
	return &Client{}, nil
}

// GetValues queries the environment for keys
func (c *Client) GetValues(keys []string) (map[string]string, error) {
	vars := make(map[string]string)
	// http://169.254.254.254/latest/user-data
	return vars, nil
}

func (c *Client) WatchPrefix(prefix string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	<-stopChan
	return 0, nil
}
