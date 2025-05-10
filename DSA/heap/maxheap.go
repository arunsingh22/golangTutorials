package main

type maxHeap []string

// 5 container/heap interface func must be implemented
func (h maxHeap) Less(i, j int) bool {
	return len(h[i]) >= len(h[j])
}

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(string))
}

func (h *maxHeap) Pop() any {
	n := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return n
}

func (h maxHeap) Peek() any {
	return h[0]
}
