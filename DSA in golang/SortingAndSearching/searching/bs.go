package main

import "slices"


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



func main() {
	arr := []int{11, 2, -3, 4, 5, 6, 7, 10, 16, 20}

	low := 0
	high := len(arr) - 1
	mid := 0
	target := 10

	// binary search code.
	for low <= high {
		mid = low + (high-low)/2
		if arr[mid] == target {
			fmt.Printf("Target Found: at %v", mid)
			break
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}