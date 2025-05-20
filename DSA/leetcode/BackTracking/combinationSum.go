package main

import "slices"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	comb := []int{}
	backTrack(0, candidates, target, comb, &result)
	return result
}

func backTrack(idx int, candidates []int, target int, comb []int, result *[][]int) {
	// Base case
	if idx == len(candidates) {
		if target == 0 {
			f := make([]int, len(comb))
			copy(f, comb)
			*result = append(*result, f)
		}
		return
	}
	if candidates[idx] <= target {
		// pick
		comb = append(comb, candidates[idx])
		backTrack(idx, candidates, target-candidates[idx], comb, result)
		comb = comb[:len(comb)-1] // remove the last element
	}
	// not-pick
	backTrack(idx+1, candidates, target, comb, result)
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	comb := []int{}
	slices.Sort(candidates)
	backTrack2(0, candidates, target, comb, &result)
	return result
}

func backTrack2(idx int, candidates []int, target int, comb []int, result *[][]int) {
	// Base case
	if target == 0 {
		f := make([]int, len(comb))
		copy(f, comb)
		*result = append(*result, f)
		return
	}

	for i := idx; i < len(candidates); i++ {
		if i > idx && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] <= target {
			// pick
			comb = append(comb, candidates[i])
			backTrack2(i+1, candidates, target-candidates[i], comb, result)
			comb = comb[:len(comb)-1] // remove the last element
		}
	}
}
