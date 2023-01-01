package main

import (
	"fmt"
	"sync"
)

func main() {
	a := asChan(1, 2, 3, 4, 5)
	b := asChan(6, 7, 8, 9, 10)
	c := asChan(16, 17, 18, 19, 20)
	d := merge(a, b, c)

	for val := range d {
		fmt.Printf("%d\n", val)
	}
}

func merge(NChans ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(NChans))
		defer close(out) //commenting this yeilds a deadlock situation but uncommenting gives nothing on screen as the go func finishes before the below for loop executues.
		for _, ch := range NChans {
			// for every channel ch we will have a seperate goroutine.
			go func(c <-chan int) {
				for val := range c {
					out <- val
				}
				wg.Done()
			}(ch)
		}
		wg.Wait()
	}()
	return out
}

func asChan(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, x := range nums {
			ch <- x
		}
		close(ch)
	}()
	return ch
}
