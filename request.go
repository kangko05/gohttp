package gohttp

import (
	"context"
	"io"
	"net/http"
)

var HEADER = map[string]string{
	"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
	"Accept-Language":           "ko-KR,ko;q=0.9,en-US;q=0.8,en;q=0.7",
	"Accept-Encoding":           "gzip, deflate, br",
	"Connection":                "keep-alive",
	"Cache-Control":             "max-age=0",
	"Sec-Ch-Ua":                 "\"Google Chrome\";v=\"122\", \"Chromium\";v=\"122\", \"Not-A.Brand\";v=\"99\"",
	"Sec-Ch-Ua-Mobile":          "?0",
	"Sec-Ch-Ua-Platform":        "\"Windows\"",
	"Sec-Fetch-Dest":            "document",
	"Sec-Fetch-Mode":            "navigate",
	"Sec-Fetch-Site":            "none",
	"Sec-Fetch-User":            "?1",
	"Upgrade-Insecure-Requests": "1",
	"Referer":                   "https://www.google.com/",
}

func newRequest(ctx context.Context, method, url string, body io.Reader, referer ...string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	for key, val := range HEADER {
		req.Header.Set(key, val)
	}

	if len(referer) > 0 {
		req.Header.Set("Referer", referer[0])
	}

	return req, nil
}
