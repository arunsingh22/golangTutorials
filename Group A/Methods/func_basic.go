package main

import (
	"fmt"
	"strconv"
)

func foo() int {
	fmt.Println("foo")
	return -1
}
func worker(x int) string {
	return strconv.Itoa(x) + "worker"
}
func takesfunc(callback func(int) string) string {
	return callback(5)
}
func main() {
	x := foo
	fmt.Printf("%T\t%T\n", x, foo)
	// x()
	fmt.

	//Anonymous function
	y := func(x int) int {
		return x * -1
	}(5)
	fmt.Println(y)

	//callback functions
	fmt.Println(takesfunc(worker))
}
