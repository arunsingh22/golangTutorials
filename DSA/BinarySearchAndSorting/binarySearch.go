package main

import (
	"fmt"
	"slices"
)

// slices.BinarySearch(arr, target) (int, bool)
// slices.BinarySearchFunc(arr, target, func(a, b T) int) int
//  Reasoning: this is used to find the correct index/location for
//             a new element in a sorted array where the logic of being sorted is customly defined
// slices.Index(arr, target) int // lowerbound
// slices.LastIndex(arr, target) int // upperbound

func binarysearch(arr []int, target int) int {
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

// customeBS := slices.BinarySearchFunc(arr, target, func(a, b Employee) int {
// 	if len(a.Name) > len(b.Name) {
// 		return 1 // a should come after
// 	}
// 	 if len(a.Name) < len(b.Name) {
// 		return -1 // b should come before
// 	}
// 	return 0 // both are equal
// })

func lb(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			result = mid
			high = mid - 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

func ub(nums []int, target int) int {
	low, high := 0, len(nums)-1
	result := -1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			result = mid
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}
func main() {
	lowerBound()
	fmt.Println("---------------")
	upperBound()
	// arr := []int{11, 2, -3, 4, 5, 6, 7, 10, 16, 20}

	// low := 0
	// high := len(arr) - 1
	// mid := 0
	// target := 10

	// // binary search code.
	// for low <= high {
	// 	mid = low + (high-low)/2
	// 	if arr[mid] == target {
	// 		fmt.Printf("Target Found: at %v", mid)
	// 		break
	// 	} else if arr[mid] < target {
	// 		low = mid + 1
	// 	} else {
	// 		high = mid - 1
	// 	}
	// }
}
