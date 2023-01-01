package main

import (
	"fmt"
	"sort"
)

func bs(nums []int) {
	low := 0
	high := len(nums) - 1
	target := 9

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			fmt.Println("Present at: ", mid)
			break
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}

func lower_bound(nums []int, target int) {
	low := 0
	high := len(nums) - 1
	ans := -1

	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			ans = mid //potential ans
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Println("Lower_bound", ans)
}

func sortIncreasing(s []int) {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 3, 4, 5} //
	// bs(nums)
	sortIncreasing(nums)

	for i, v := range nums {
		fmt.Printf("%d : %d\n", i, v)
	}

	// char := 'a' //char are represented as int32 internally
	// fmt.Printf("%T,%v", char, char)

	// target := 20
	// lower_bound(nums, target)
	// index := sort.Search(len(nums), func(i int) bool {
	// 	return nums[i] >= target
	// })
	// fmt.Println(index)
	// if index < len(nums) && nums[index] == target {
	// 	fmt.Println("Found:", index)
	// } else {
	// 	fmt.Println("Not present")
	// }

}
