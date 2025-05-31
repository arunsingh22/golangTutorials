package main

import "math"

// https://www.youtube.com/watch?v=hnswaLJvr6g
func maxProduct(nums []int) int {
	pref, suff := 1, 1
	ans := int(math.Inf(-1)) // INT_MIN equivalent

	n := len(nums)
	for i := 0; i < n; i++ {
		if pref == 0 {
			pref = 1
		}
		if suff == 0 {
			suff = 1
		}
		pref *= nums[i]
		suff *= nums[n-i-1]
		ans = max(ans, max(pref, suff))
	}
	return ans
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
