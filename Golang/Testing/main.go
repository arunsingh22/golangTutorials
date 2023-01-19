package main

import (
	"fmt"
)

func Sum(input ...int) int {
	sum := 0
	for _, r := range input {
		// fmt.Println(idx, r)
		sum += r
	}
	return sum
}
func main() {
	fmt.Println(Sum(1, 2, 3, 4, 5))
}
