// https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/
// For C++, Java,GO: We need to custom our comparator since the default comparator is string comparator which compares by lexicographically order, it means for example: "123" < "14".

func main() {

}
func kthLargestNumber(nums []string, k int) string {
	minheap := MinHeap(nums)
	heap.Init(&minheap)
	for minheap.Len() > k {
		heap.Pop(&minheap)
	}
	return heap.Pop(&minheap).(string)
}

// sort.Interface
type MinHeap []string

// type any interface{}

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Less(i, j int) bool {
	if len(h[i]) == len(h[j]) {
		return h[i] < h[j]
	}
	return len(h[i]) < len(h[j]) // string which is smaller in len will be before.
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}
func (h *MinHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}

