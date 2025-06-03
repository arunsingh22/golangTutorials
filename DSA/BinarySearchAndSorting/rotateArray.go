package main

import "slices"

// https://www.youtube.com/watch?v=BHr381Guz3Y
func rotate(nums []int, k int) {
	// Solution 1: TC: O(N) and SC: O(N)
	N := len(nums)
	tmp := make([]int, N)
	for i := 0; i < len(nums); i++ {
		tmp[(i+k)%N] = nums[i]
	}
	copy(nums, tmp)

	// Solution 2: TC: O(N) and SC: O(1)
	// If the mod is zero it mean no rotation is required this is a smart optimization done to
	// prevent TTL for large K.
	if k%N == 0 {
		return
	}
	// These are edge case if k > N , this has manily no effect as the multiple of k yeilds same rotation
	// therefore we trim down k by taking mod with N, this is also important as we are slicing and it prevents
	// out of bound slicing.
	k = k % N
	if N > 1 {
		slices.Reverse(nums)
		slices.Reverse(nums[:k])
		slices.Reverse(nums[k:])
	}

}
