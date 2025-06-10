package gohttp

import "time"

type Option func(*Client)

func WithRetries(n int) Option {
	if n > 0 {
		return func(c *Client) {
			c.retries = n
		}
	}

	return nil
}

func WithTimeout(dur time.Duration) Option {
	if dur > 0 {
		return func(c *Client) {
			c.timeout = dur
		}
	}

	return nil
}
