package main

import (
	"fmt"
	"reflect"
	"time"
)

// done chan is a bool channel which is analogous to ctx.Done() it is used for graceful shutdown
// of goroutines.
func fib(c chan int, done chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("Time to stop fib func")
			return
		case v := <-c:
			fmt.Println(v * 2)
		}
	}
}

func main() {
	c := make(chan int)
	done := make(chan struct{})
	fmt.Println("done chan size: ", reflect.TypeOf(done))
	go fib(c, done)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- struct{}{}
	}()
	time.Sleep(10 * time.Second)
}
