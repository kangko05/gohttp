package gohttp

import (
	"context"
	"time"
)

type RateLimiter struct {
	ctx        context.Context
	maxTokens  int
	tokensChan chan struct{}
}

// this will automatically start the process as goroutine
func NewRateLimiter(ctx context.Context, maxTokens int) *RateLimiter {
	rl := &RateLimiter{
		ctx:        ctx,
		maxTokens:  maxTokens,
		tokensChan: make(chan struct{}, maxTokens),
	}

	for range rl.maxTokens {
		rl.tokensChan <- struct{}{}
	}

	go rl.start()

	return rl
}

func (rl *RateLimiter) start() {
	ticker := time.NewTicker((time.Second / time.Duration(rl.maxTokens)))

	for {
		select {
		case <-rl.ctx.Done():
			// cleanup maybe
			ticker.Stop()
			close(rl.tokensChan)
			return
		case <-ticker.C:
			rl.refillTokens()
		}
	}
}

func (rl *RateLimiter) refillTokens() {
	if len(rl.tokensChan) < rl.maxTokens {
		rl.tokensChan <- struct{}{}
	}
}

func (rl *RateLimiter) GetToken() {
	<-rl.tokensChan
}
