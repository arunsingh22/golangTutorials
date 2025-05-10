package main

import "fmt"

func main() {
	// nums := []int{3, 2, 1, 5, 6, 4}
	// minheap := MinHeap(nums) //
	// heap.Init(&minheap)

	// for minheap.Len() > 0 {
	// 	x := heap.Pop(&minheap) // use heap.Pop to get the min element
	// 	fmt.Println(x)
	// }

	// names := []string{"ar", "arun", "s", "shilpayadav", "arunsinghisgoodboy"}
	// maxheap := maxHeap(names)
	// heap.Init(&maxheap)

	// x := maxheap.Peek()
	// fmt.Println(x.(string))
	// heap.Pop(&maxheap)
	// x = maxheap.Peek()
	// fmt.Println(x.(string))

	// for maxheap.Len() > 0 {
	// 	x := maxheap.Peek()
	// 	fmt.Println(x.(string))
	// }

	nums := []string{"16744716392823", "243637189", "492312", "1432609310668", "4356708604", "5545", "45020005347491", "34755411082555", "9953439611291093371", "995", "8382023684", "657161768802942409", "46", "9812786866943202", "4908788443448123083", "1242403381", "6", "65184269450053424089", "7498", "8354396", "83652", "2727223972404226", "41216376182091660100", "41429", "1", "9513", "89724", "82308481001", "518620303662543617", "231451", "61510922685926249451", "854102252", "8711964215226", "676769523668", "9878836196893135", "4"}
	fmt.Println(kthLargestNumberUsingHeap(nums, 2))
}
