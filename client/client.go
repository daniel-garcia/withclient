package client

import (
	"golang.org/x/net/context"
)

type Client struct {
	ctx context.Context
}

func NewClient() *Client {
	return &Client{
		ctx: context.Background(),
	}
}

func (c *Client) WithValue(key, value interface{}) *Client {
	return &Client{
		ctx: context.WithValue(c.ctx, key, value),
	}
}

func (c *Client) Foo() int {
	fooVal := c.ctx.Value("foo")
	if fooVal != nil {
		foo, ok := fooVal.(int)
		if ok {
			return foo
		}
	}
	return 42
}

func (c *Client) Bar() string {
	barVal := c.ctx.Value("bar")
	if barVal != nil {
		bar, ok := barVal.(string)
		if ok {
			return bar
		}
	}
	return "baz"
}
