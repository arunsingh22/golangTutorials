package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1
	b, ok := <-ch      //reading from chan
	fmt.Println(b, ok) //1, true

	close(ch) //closing the chan

	c, ok := <-ch
	fmt.Println(c, ok) // 0,false

	// NOTE: If the chan is unbuffered then it will cause a deadlock
}
