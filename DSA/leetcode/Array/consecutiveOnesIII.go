package main

// https://leetcode.com/problems/max-consecutive-ones-iii/description/
// Sliding window
func consecutiveOnesIII(nums []int, k int) int {
	ans := 0
	zeroCnt := 0
	i, j := 0, 0

	for j < len(nums) {
		if nums[j] == 0 {
			zeroCnt++
		}
		if zeroCnt <= k {
			// take that subarray len
			ans = max(ans, j-i+1)
		}
		for zeroCnt > k {
			// zeroCnt has exceed k ie move i now
			if nums[i] == 0 {
				zeroCnt--
			}
			i++
		}
		j++
	}
	return ans
}
