package gohttp

import (
	"context"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	ctx     context.Context
	cl      *http.Client
	timeout time.Duration
	retries int
}

func NewClient(ctx context.Context) *Client {
	jar, _ := cookiejar.New(nil)

	cl := &http.Client{
		Jar: jar,
	}

	client := &Client{
		ctx:     ctx,
		cl:      cl,
		timeout: 30 * time.Second,
		retries: 3,
	}

	return client
}

func (c Client) Context() context.Context {
	return c.ctx
}
