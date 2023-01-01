package main

import "fmt"

// what is the value returned when a channel is close in ok idiom.
// Ans:  https://www.educative.io/answers/what-is-closec-chan---type-in-golang

// How to check if a channel is closed or not

// Who should close the channels?
// Ans: Senders close channels and receivers only check

// "Is there any other method to create a channel without the make function?" -- No. Channel types do not support composite literals and var c chan t only declares the variable, it does not initialize the channel, leaving it nil Refer this: https://stackoverflow.com/questions/71626679/can-i-create-channel-without-using-the-make-function#:~:text=%22Is%20there%20any%20other%20method,the%20channel%2C%20leaving%20it%20nil%20.

func main() {
	example := make(chan string, 10)

	example <- "arun"
	example <- "singh"
	fmt.Println(<-example)
	close(example)
	fmt.Println(<-example)
	result, ok := <-example

	fmt.Println(result)
	fmt.Println(ok)

}
