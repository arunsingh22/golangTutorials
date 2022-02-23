package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 5
	fmt.Println(<-c)
}
