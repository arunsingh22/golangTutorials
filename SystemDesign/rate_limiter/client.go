package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel() // Ensure the leakProcess goroutine can be stopped

	usr1LB := NewLeakyBucket(ctx, 5, 2)
	// usr2LB := NewLeakyBucket(context.Background(), 10, 3)
	// fmt.Println(usr1LB, usr2LB)

	for range 10 {
		go usr1LB.IsRequestAllowed("1")
	}
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
