package main

import (
	"fmt"
	"sort"
)

// Refer : https://go.dev/src/sort/search.go

// For instance, given a slice data sorted in ascending order,
// the call Search(len(data), func(i int) bool { return data[i] >= 23 })
// returns the smallest index i such that data[i] >= 23. If the caller
// wants to find whether 23 is in the slice, it must test data[i] == 23
// separately.
//
// Searching data sorted in descending order would use the <=
// operator instead of the >= operator.

// Note: sort.Search does both binary search and lower_bound search when duplicates present.

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 5, 6, 6, 9, 10}
	x := 1

	// Binary search
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= x // >= is mandatory cannot use ==
	})
	if i < len(nums) && nums[i] == x {
		// x is present at nums[i]
		fmt.Println(i)
	} else {
		// x is not present in nums,
		// but i is the index where it would be inserted.
		fmt.Println("Not present: ", i)
	}

	// char := 'a'
	// fmt.Printf("%T,%v", char, char)
}
