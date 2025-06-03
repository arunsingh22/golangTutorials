package main

// https://www.youtube.com/watch?v=Vzs_vlCIFEw
func checkArrayIsSortedAndRotated(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	cnt := 1
	// we are kind of simulating the sliding window where we are now trying to find out
	// if there exits a subarray within the 2N array size which has a sorted seq
	// example : [3,4,5,1,2] --> [3,4,5,1,2|3,4,5,1,2]
	for i := 1; i < 2*n; i++ {
		if nums[(i-1)%n] <= nums[i%n] {
			cnt++
		} else {
			cnt = 1
		}
		// if anytime we find a subarray which is equal to N
		// this means the array is sorted and rotated.
		if cnt == n {
			return true
		}
	}
	return false
}
