package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	ch1 := make(chan int)
	timer := time.NewTimer(5 * time.Second)   //ONE TIME TIMER
	ticker := time.NewTicker(1 * time.Second) // REPEATED TIMER.

	// Select is like a switch case which works with concurrent channels
	// select {
	// case <-ticker.C:
	// 	fmt.Println("tick")
	// case <-timer.C:
	// 	fmt.Println("Timer has fired after 5 seconds!")
	// }

	defer timer.Stop()
	defer ticker.Stop()

	go func() {
		ch1 <- 100
		close(ch1)
	}()

	// When we use a for loop together with a select statement,
	// we create a method for continually checking multiple channels:
	for {
		select {
		case x, ok := <-ch1:
			if ok {
				fmt.Println("channel closed", x)
				// set the ch1 to nil as nil channels are removed from the select block as
				// reading from them always block and select knows this
				// this is great way to soft delete a channel
				ch1 = nil
			}
		case <-ticker.C:
			fmt.Println("tick", reflect.TypeOf(ch1))
		case <-timer.C:
			fmt.Println("Timer has fired after 5 seconds!")
			return
		}
	}
}
