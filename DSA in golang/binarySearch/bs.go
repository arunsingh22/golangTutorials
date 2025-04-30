package binarysearch

import "slices"

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




func rawbinarysearch()

func binarysearch(arr []int, target int) int {
	// len of arr is 0 return -1

	idx, ok := slices.BinarySearch(arr, target)
	if ok {
		return idx
	}
	return -1
}

type Employee struct {
	Name string
	Age  int
}

binarysearch := slices.BinarySearchFunc(arr, target, func(a, b Employee) int {
	if len(a.Name) > len(b.Name) {
		return 1
	} 
	 if len(a.Name) < len(b.Name) {
		return -1
	} 
	return 0 
})
