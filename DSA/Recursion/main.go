package main

import "fmt"

func main() {
	// printingNtoOne(5)
	printingOneToN(5)
}

// head recursion: post evaluation
func printingOneToN(n int) {
	// BC
	if n == 1 {
		fmt.Println(n)
		return
	}
	printingOneToN(n - 1)
	fmt.Println(n)
}

// tail recursion
func printingNtoOne(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printingNtoOne(n - 1)
}
