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

func NewTokenBucker(ctx context.Context, capacity int, fillRate int) IRateLimiter {
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
	if tb.tokens >= tb.capacity {
		fmt.Println("bucket is full, refill halted!")
	}
	ticker := time.NewTicker(time.Duration(fillrate) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:

		case <-ctx.Done():
			fmt.Println("exiting refillBucket routine..")
			return
		}
	}
}

func (tb *TokenBucket) IsRequestAllowed(req any) bool {
	return true
}
