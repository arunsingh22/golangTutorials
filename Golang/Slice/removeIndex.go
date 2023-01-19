package main

import "fmt"

// Output : [1 3 4 4] Not so proper way of slicing the array.
func rmidx(s []int, idx int) []int {
	return append(s[:idx], s[idx+1:]...)
}

// [1 3 4]
func properrmidx(s []int, idx int) []int {
	newSlice := []int{}
	newSlice = append(newSlice, s[:idx]...)
	return append(newSlice, s[idx+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println("Before: ", slice)
	newSlice := properrmidx(slice, 1)
	fmt.Println("After:", newSlice)
}
