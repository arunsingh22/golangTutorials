package main

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// T: O(log n)
// S: O(log n)
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// T: O(log n) - because when we remove the root (min) we need to reorder.
// S: O(log n)
func (h *MinHeap) Pop() any {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}
