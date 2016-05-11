package other

import (
	"github.com/daniel-garcia/withcontext/client"
)

type Other struct {
	client BarClient
}

type BarClient interface {
	Bar() string
	WithValue(key, value interface{}) BarClient
}

type ContextToBarClientAdapter struct {
	*client.Client
}

func NewContextToBarClientAdapter(c *client.Client) BarClient {
	return &ContextToBarClientAdapter{c}
}

func (c *ContextToBarClientAdapter) WithValue(key, value interface{}) BarClient {
	return &ContextToBarClientAdapter{
		c.Client.WithValue(key, value),
	}
}

func NewOther(c BarClient) *Other {
	return &Other{
		client: c,
	}
}
