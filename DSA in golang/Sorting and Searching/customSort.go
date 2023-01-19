package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	a, b int
}

func main() {
	nums := []int{1, 5, 6, 2, 10, 4, 7}
	target := 6

	res := []int{}
	s := []Pair{}

	for i, k := range nums {
		s = append(s, Pair{i, k})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i].b < s[j].b
	})

	//while loop
	for i, j := 0, len(s)-1; i < j; {
		if s[i].b+s[j].b == target {
			res = append(res, s[i].a, s[j].a)
			break
		} else if s[i].b+s[j].b > target {
			j--
		} else {
			i++
		}
	}
	fmt.Println(res)

}
