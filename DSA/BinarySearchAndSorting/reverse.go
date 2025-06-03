package main

import (
	"fmt"
)

// In Golang, if you have multiple conditional checks in a for loop, you combine them using
//  logical operators like && (AND), || (OR), etc.

// But important point in Go:
// - The init section allows multiple variable declarations separated by commas (e.g., i, j := 0, 10).
// - The condition is a single boolean expression (e.g., i < j).
// - The post statement allows multiple assignments, but you must use i, j = newI, newJ syntax.
// - You cannot separate multiple post statements with commas like in C/C++.
// for init1, init2 := val1, val2; (condition1) && (condition2); post1, post2 = update1, update2 {
//     // loop body
// }
// for i, j := 0, 10; (i < j) && (i < 5); i, j = i+1, j-1 {
//     fmt.Println(i, j)
// }

func maiN() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	// Most elegant way to do in-place reverse
	// Go version of while loop.
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	fmt.Println(nums)

	// reverse a string chars as strings doesn't has any inbuilt func
	name := "arun singh" // this readOnly and cannot be modified

	r := []rune(name) // creating a new slice from name

	// in-place reverse
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	reverseString := string(r)
	fmt.Println(reverseString) // this is a new string
	fmt.Println(&name, &reverseString)
}
