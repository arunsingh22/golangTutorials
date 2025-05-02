package main

import (
	"fmt"
	"slices"
)

// NOTE: The slices package itself has both Sorting and search funcs in-built
// Sort string characters? → slices.Sort([]rune)
// slices.Sort(arr)
// slices.Contains(arr, target) bool
// slices.IsSorted(arr) bool
// slices.Reverse(arr)
// slices.Max(arr)
// slices.Min(arr)
// slices.BinarySearch(arr, target) (int, bool)
// slices.BinarySearchFunc(arr, target, func(a, b T) int) int
// slices.Index(arr, target) int // lowerbound
// slices.LastIndex(arr, target) int // upperbound

func main() {
	// sorting array of integers.
	s := []int{4, 2, 3, 1}
	slices.Sort(s) // default is sorting in ascending order
	fmt.Println(s)

	slices.SortStableFunc(s, func(a, b int) int {
		return a - b // same as ascending order
	})
	fmt.Println("Sorting after custome sorting: ", s)

	//Sorting a array of strings
	x := []string{"arun", "s", "yadav", "aaaa"}
	slices.Sort(x) // default is sorting in ascending order
	fmt.Println(x)

	// sorting in descending order
	// 	The comparator tells Go how to order two elements:
	// Return < 0 → a should come before b
	// Return 0 → a and b are equal
	// Return > 0 → a should come after b
	slices.SortFunc(x, func(a, b string) int {
		if len(a) < len(b) {
			return 1 // +ve indicates a will come after b
		}
		if len(a) > len(b) {
			return -1 // a -ve indicates that a will come before b
		}
		return 0 // both a and b are considered equal
	})
	fmt.Println("After sorting based on str len: ", x)

	d := []rune{'e', 'f', 'z', 'a', 'b'}
	slices.Sort(d)
	fmt.Println([]rune(d)) // default is sorting in ascending order

	// slices.Reverse(x)
	// fmt.Println("After reversing: ", x)

}
