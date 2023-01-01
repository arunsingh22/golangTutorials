package main

import (
	"fmt"
	"time"
)

// A generator is just a func which creates a channel and also populates it with data
// using a goroutine within itself, there it's self contained of producing both the channel and it's values

func generator() <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			out <- i * 2
			// Imagine the generator is doing somthhing very complex operation to produce the data.
			time.Sleep(2 * time.Millisecond)
		}
	}()
	return out
}
func consumer(ch <-chan int) {
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				ch = nil
			}
			fmt.Println(v)
		}
	}
}
func main() {

	c := generator()
	go consumer(c)
	time.Sleep(2 * time.Second)
	fmt.Println("Main exit")

}
