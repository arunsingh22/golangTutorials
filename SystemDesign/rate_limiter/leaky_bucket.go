package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Implementaion of leaky bucket
// https://medium.com/@oneconfusedindian/decoding-complexity-leaky-bucket-algorithm-in-golang-59d522981b95

type IRateLimiter interface {
	IsRequestAllowed(req any) bool
}

type LeakyBucket struct {
	bucket     []any // hyperparameter
	capacity   int
	outputRate int // hyperparameter
	mu         *sync.Mutex
}

// This is part of the interface
func (lb *LeakyBucket) IsRequestAllowed(req any) bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	if len(lb.bucket) >= lb.capacity {
		// drop the request as we cannot add into the queue
		fmt.Println("Requeset dropped!")
		return false
	}
	lb.bucket = append(lb.bucket, req)
	fmt.Println("Requeset Allowed!")
	return true
}

func (lb *LeakyBucket) removeFromBucket(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if len(lb.bucket) > int(lb.outputRate) {
				lb.mu.Lock()
				lb.bucket = lb.bucket[lb.outputRate:]
				lb.mu.Unlock()
			}
		case <-ctx.Done():
			fmt.Println("leaky bucket has stopped!")
			return
		}
	}
}

func NewLeakyBucket(ctx context.Context, capcity int, outputRate int) IRateLimiter {
	if capcity <= 0 {
		panic("capacity must be positive")
	}
	if outputRate <= 0 {
		panic("outputRate must be positive")
	}
	lb := &LeakyBucket{
		bucket:     []any{}, // pre-allocate the slices
		capacity:   capcity,
		outputRate: outputRate,
		mu:         &sync.Mutex{},
	}
	go lb.removeFromBucket(ctx)
	return lb
}
