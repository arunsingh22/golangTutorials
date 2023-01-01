package main

import (
	"fmt"
)

func main() {
	// working code.
	// c := make(chan int, 1) //buffered chan
	// c <- 1
	// fmt.Println("Reading from channel: ", <-c)

	// Working code.
	// c := make(chan int) //unbuffered chan
	// go func(c chan<- int) {
	// 	c <- 2 //writing into chan
	// }(c)
	// fmt.Println("Reading from channel: ", <-c) //getting blocked here for other goroutines to write.

	// DeadLock code.
	c := make(chan int) //unbuffered chan
	c <- 1              //getting blocked here!!!

	// This never gets executed aka deadcode
	go func(c chan int) {
		fmt.Println("Reading from channel: ", <-c)
	}(c)

}
