// https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
// For C++, Java,GO: We need to custom our comparator since the default comparator is string comparator which compares by lexicographically order, it means for example: "123" < "14".
package main

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"
)

// heap.Interface
type ownMaxHeap []string

func (h ownMaxHeap) Len() int {
	return len(h)
}

func (h ownMaxHeap) Less(i, j int) bool {
	x, _ := strconv.Atoi(h[i])
	y, _ := strconv.Atoi(h[j])
	return x > y
}

// check this
func (h ownMaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *ownMaxHeap) Push(x any) {
	*h = append(*h, x.(string))
}

func (h *ownMaxHeap) Pop() any {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

// method 2
func kthLargestNumberUsingHeap(nums []string, k int) string {
	var result string
	maxHeap := ownMaxHeap(nums)
	heap.Init(&maxHeap)

	for i := 0; i != k; i++ {
		result = heap.Pop(&maxHeap).(string)
		fmt.Println(k, result)
	}
	return result
}

// method 1:
func kthLargestNumberUsingSort(nums []string, k int) string {
	slices.SortFunc(nums, func(i, j string) int {
		a, _ := strconv.Atoi(j)
		b, _ := strconv.Atoi(i)
		return a - b
	})
	// fmt.Println(nums)
	return nums[k-1]
}
