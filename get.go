package gohttp

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func (c *Client) get(url string, referer ...string) (*Response, error) {
	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()

	req, err := newRequest(ctx, "GET", url, nil, referer...)
	if err != nil {
		return nil, err
	}

	resp, err := c.cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rb, err := decodeResponseBody(resp)
	if err != nil {
		return nil, err
	}

	return &Response{
		Status: resp.StatusCode,
		Header: resp.Header.Clone(),
		Body:   rb,
	}, nil
}

func (c *Client) Get(url string, referer ...string) (*Response, error) {
	var lastErr error

	for i := range c.retries {
		if i > 0 {
			dur := time.Duration(math.Pow(2, float64(i))) * time.Second
			sleepdur := dur + ((time.Duration(rand.IntN(1000))) * time.Millisecond)

			time.Sleep(sleepdur)
		}

		resp, err := c.get(url, referer...)
		if err != nil {
			lastErr = fmt.Errorf("req failed after %d attempts: %v", i+1, err)
			continue
		}

		return resp, nil
	}

	return nil, lastErr
}
