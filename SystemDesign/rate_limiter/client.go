package main

import (
	"context"
	"fmt"
	"time"
)

type RateLimiterType string

const (
	TokenBucketType RateLimiterType = "TOKEN_BUCKET"
	LeakyBucketType RateLimiterType = "LEAKY_BUCKET"
	FixedWindowType RateLimiterType = "FIXED_WINDOW"
)

// Context
type RateLimiter struct {
	rateLimiter IRateLimiter
}

// Factory Method
func (r *RateLimiter) GetRateLimiter() IRateLimiter {
	return r.rateLimiter
}

func (r *RateLimiter) SetStrategy(rl IRateLimiter) {
	r.rateLimiter = rl
}

func (r *RateLimiter) Execute(req any) bool {
	if r.rateLimiter == nil {
		fmt.Println("No rate limiter set!")
		return false
	}
	return r.rateLimiter.IsRequestAllowed(req)
}

func RateLimiterFactory(ctx context.Context, limiterType RateLimiterType, capacity, rate int) IRateLimiter {
	switch limiterType {
	case TokenBucketType:
		return NewTokenBucket(ctx, capacity, rate)
	case LeakyBucketType:
		return NewLeakyBucket(ctx, capacity, rate)
	case FixedWindowType:
		return NewLeakyBucket(ctx, capacity, rate)
	default:
		panic("Unsupported rate limiter type")
	}
}

// Problem Statement
// a. Functional Requirements
// The system should allow defining rate limits per user. A rate limit consists of:
// User ID (identifier for the user) for each user it should have **
// Maximum Requests (number of allowed requests).
// Time Window (duration in seconds).
// Rate-Limiting Algorithm (e.g., Counter, Token Bucket, Leaky Bucket).
// b. Non-functional Requirements
// - Extensible
// - Performance
// - Concurrecny
// c. Example
// Identifying Key Classes
// Class Diagram
// Implementation

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure the leakProcess goroutine can be stopped

	usr1LB := RateLimiterFactory(ctx, LeakyBucketType, 5, 2)
	// usr2LB := NewLeakyBucket(context.Background(), 10, 3)
	// fmt.Println(usr1LB, usr2LB)

	// usrTB := NewTokenBucket(ctx, 13, 2)

	r := RateLimiter{}
	r.SetStrategy(usr1LB)
	for range 10 {
		go r.Execute("1")
	}

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Done")

}
