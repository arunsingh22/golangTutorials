package main

// https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308

import (
	"fmt"
	"sync"
)

func main() {
	a := asChan(1, 2, 3, 4, 5)
	b := asChan(6, 7, 8, 9)
	c := asChan(10, 11, 12, 13, 14)

	d := merge(a, b, c)
	for val := range d {
		fmt.Printf("%d,", val)
	}
}

// We are creating N goroutines per Channel and every goroutine adds into the out channel.
func merge(NChans ...<-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(NChans))
		defer close(outCh)
		for _, ch := range NChans {
			go func(ch <-chan int) {
				for val := range ch {
					outCh <- val
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
		// close(outCh)
	}()
	return outCh
}

func asChan(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, val := range nums {
			out <- val
		}
	}()
	return out
}
