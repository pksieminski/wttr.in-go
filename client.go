package wttr

import "net/http"

type Client struct {
	config Configuration
	client http.Client
}

func NewClient() (*Client, error) {
	c, err := NewConfiguration()

	if err != nil {
		return nil, err
	}

	return NewClientWithConfiguration(c), nil
}

func NewClientWithConfiguration(c *Configuration) *Client {
	return &Client{
		config: *c,
		client: http.Client{},
	}
}
