package main

import (
	"fmt"
	"sort"
)

func main() {
	//Increasing order sorting.
	s := []int{4, 2, 3, 1}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)

	//Sorting in decreasing order.
	x := []string{"arun", "s", "yadav", "aaaa"}
	sort.Slice(x, func(i, j int) bool {
		return x[i] > x[j]
	})
	fmt.Println(x)
}
