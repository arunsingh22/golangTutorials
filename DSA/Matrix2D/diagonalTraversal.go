package main

import (
	"fmt"
	"slices"
)

// Even though Go doesn't have a built-in SortedMap,
// you can achieve the same functionality using a combination of Go's existing features:

// Using a Separate Slice for Keys:
// This is the most common approach, and the one used in the corrected code I provided.

// Store your key-value pairs in a regular Go map.
// Maintain a separate slice of keys.
// When you need to iterate in sorted order:
// Sort the slice of keys.
// Iterate over the sorted slice, looking up the corresponding values in the map.

func findDiagonalOrder(mat [][]int) []int {
	m := make(map[int][]int) // init map with no len
	result := []int{}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if _, ok := m[i+j]; ok {
				// append the result
				m[i+j] = append(m[i+j], mat[i][j])
			} else {
				m[i+j] = []int{mat[i][j]}
			}
		}
	}
	fmt.Println(m)

	// build the result set
	for k, val := range m {
		if k&1 != 0 {
			result = append(result, val...)
		} else {
			// even => reverse order
			slices.Reverse(val)
			result = append(result, val...)
		}
	}
	fmt.Println(result)
	return result
}
