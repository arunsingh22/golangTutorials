package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// https://www.educative.io/answers/how-to-use-string-builder-in-golang
func main() {
	nums := []int{12, 0, 5, 4}
	fmt.Println(largestNumber(nums))
}

// using custom sorting
func largestNumber(nums []int) string {

	strNums := make([]string, len(nums))
	for i, n := range nums {
		strNums[i] = strconv.Itoa(n) // converting number to ASCII
	}
	sort.Sort(ByLargestNumber(strNums))

	var str strings.Builder // 1st step
	for _, n := range strNums {
		str.WriteString(n) // 2nd step
	}
	ans := str.String() // 3rd step , that's it
	if ans[0] == '0' {
		return "0"
	}
	return ans
}

// sort.Interface{}
type ByLargestNumber []string

func (n ByLargestNumber) Len() int           { return len(n) }
func (n ByLargestNumber) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByLargestNumber) Less(i, j int) bool { return n[i]+n[j] > n[j]+n[i] }
