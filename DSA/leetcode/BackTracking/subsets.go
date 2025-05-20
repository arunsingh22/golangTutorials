package main

import "slices"

// Subset I:
func subsets(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}
	generateSubset(0, nums, subset, &result)
	return result
}

func generateSubset(idx int, nums, subset []int, result *[][]int) {
	if idx == len(nums) {
		f := make([]int, len(subset))
		copy(f, subset)
		*result = append(*result, f)
		return
	}
	subset = append(subset, nums[idx])
	generateSubset(idx+1, nums, subset, result)

	subset = subset[:len(subset)-1]
	generateSubset(idx+1, nums, subset, result)
}

// Subset II:
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	subset := []int{}
	slices.Sort(nums)
	generateSubsetWithoutDup(0, nums, subset, &result)
	return result
}

func generateSubsetWithoutDup(idx int, nums, subset []int, result *[][]int) {
	f := make([]int, len(subset))
	copy(f, subset)
	*result = append(*result, f)

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		subset = append(subset, nums[i])
		generateSubset(i+1, nums, subset, result)
		subset = subset[:len(subset)-1]
	}
}
