package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Implemention of token bucket algorithm
// https://medium.com/@oneconfusedindian/decoding-complexity-token-bucket-algorithm-ba80885acacd

type TokenBucket struct {
	ctx      context.Context
	tokens   int // hyperparameter
	capacity int
	fillRate int // hyperparameter
	mu       *sync.Mutex
}

func NewTokenBucket(ctx context.Context, capacity int, fillRate int) IRateLimiter {
	if capacity <= 0 {
		panic("capacity must be positive")
	}
	if fillRate <= 0 {
		panic("fillRate must be positive")
	}
	tb := &TokenBucket{
		ctx:      ctx,
		tokens:   0,
		capacity: capacity,
		fillRate: fillRate,
		mu:       &sync.Mutex{},
	}
	go tb.refillBucket(ctx, fillRate)
	return tb
}

func (tb *TokenBucket) refillBucket(ctx context.Context, fillrate int) {
	ticker := time.NewTicker(time.Duration(fillrate) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tb.mu.Lock()
			if tb.tokens+fillrate >= tb.capacity {
				fmt.Println("bucket is full, refill halted!")
				tb.tokens = tb.capacity
			} else {
				tb.tokens += fillrate
			}
			tb.mu.Unlock()
		case <-ctx.Done():
			fmt.Println("exiting refillBucket routine..")
			return
		}
	}
}

func (tb *TokenBucket) IsRequestAllowed(req any) bool {
	tb.mu.Lock()
	if tb.tokens >= 0 {
		tb.tokens--
		fmt.Println("Token bucket allows..")
		return true
	}
	tb.mu.Unlock()
	return false
}
