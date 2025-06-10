package gohttp

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"

	"github.com/andybalholm/brotli"
)

type Response struct {
	Status int
	Header http.Header
	Body   []byte
}

func decodeResponseBody(resp *http.Response) ([]byte, error) {
	// First, check if response body is nil
	if resp == nil || resp.Body == nil {
		return nil, fmt.Errorf("nil response or response body")
	}

	var reader io.Reader = resp.Body // Default to using response body directly

	// Fix typo in header name
	enc := resp.Header.Get("Content-Encoding")

	// Set appropriate reader based on encoding
	switch enc {
	case "gzip":
		r, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read gzip encoding: %v", err)
		}
		defer r.Close()
		reader = r
	case "br":
		reader = brotli.NewReader(resp.Body)
	case "deflate":
		r := flate.NewReader(resp.Body)
		defer r.Close()
		reader = r
	}

	// Read the entire content
	if reader == nil {
		return nil, fmt.Errorf("reader is nil after processing")
	}

	return io.ReadAll(reader)
}
