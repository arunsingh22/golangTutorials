package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	// Most elegant way to do in-place reverse
	// Go version of while loop.
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)

}
