package main

import (
	"fmt"
	"time"
)

// Fan in pattern is a technique of multiplexing multiple inputs channels into 1 or less output channels.

func generator(str string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		out <- str
	}()
	return out
}
func consumer(ch1, ch2, ch3, ch4 <-chan string) chan string {
	out := make(chan string)
	go func() {
		for {
			select {
			case v, ok := <-ch1:
				if !ok {
					ch1 = nil
				}
				out <- v
			case v, ok := <-ch2:
				if !ok {
					ch1 = nil
				}
				out <- v
			case v, ok := <-ch3:
				if !ok {
					ch1 = nil
				}
				out <- v
			case v, ok := <-ch4:
				if !ok {
					ch1 = nil
				}
				out <- v
			}
		}
	}()
	return out
}
func main() {

	c1 := generator("foo")
	c2 := generator("boo")
	c3 := generator("too")
	c4 := generator("loo")
	out := consumer(c1, c2, c3, c4)

	for x := range out {
		if x != "" {
			fmt.Println(x)
		} else {
			break
		}
	}

	// go func() {
	// 	for {
	// 		fmt.Println(<-out)
	// 		fmt.Println("Helo")
	// 	}

	// }()
	defer close(out)
	time.Sleep(5 * time.Second)
	fmt.Println("Main exit")

}
