package main

import "fmt"

func printAllSubsequence(arr []int) {
	curr := []int{}
	// generateSubseq(arr, curr, 0)
	generateDistinctSubseq(arr, curr, 0)
}

// Note: the problem here is that if the input sequence has duplicates it prints duplicate
// subsequences
func generateSubseq(arr, curr []int, idx int) {
	if idx == len(arr) {
		fmt.Println(curr)
		return
	}
	// pick the curr idx
	curr = append(curr, arr[idx])
	generateSubseq(arr, curr, idx+1)

	curr = curr[:len(curr)-1]
	generateSubseq(arr, curr, idx+1)
}

func generateDistinctSubseq(arr, curr []int, idx int) {
	if idx == len(arr) {
		fmt.Println(curr)
		return
	}

	for i := idx; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			idx++
		}
	}

	// pick the curr idx
	curr = append(curr, arr[idx])
	generateSubseq(arr, curr, idx+1)

	curr = curr[:len(curr)-1]
	generateSubseq(arr, curr, idx+1)
}

func main() {
	arr := []int{1, 2, 2}
	printAllSubsequence(arr)
}
