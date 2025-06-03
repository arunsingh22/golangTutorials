package main

// edge case
// [3,1,2,3,3,3,3,3,3]
// here standing on mid we cannot decide which side to go as the arr[low],arr[mid] and
// arr[high] all are same therefore we must simpliy drop the low and high and continue as usual
func search(nums []int, target int) bool {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		}
		// trim from both ends to remove same elements as
		// this is causing confusion which side to go
		if nums[low] == nums[mid] && nums[high] == nums[mid] {
			low++
			high--
			continue
		}
		// check if left side is sorted or not
		if nums[low] <= nums[mid] {
			if nums[low] <= target && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		// check on the right hand side
		if nums[mid] <= nums[high] {
			if nums[mid] <= target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

	}
	return false
}
