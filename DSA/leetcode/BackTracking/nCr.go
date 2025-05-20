package main

func combine(n int, k int) [][]int {
	result := [][]int{}
	comb := []int{}
	nCr(1, n, k, comb, &result)
	return result
}

func nCr(start, n, k int, comb []int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(comb))
		copy(tmp, comb)
		*result = append(*result, tmp)
		return
	}
	if start > n {
		return
	}

	// pick the current number
	comb = append(comb, start)
	nCr(start+1, n, k-1, comb, result)
	comb = comb[:len(comb)-1]
	// not pick the current number
	nCr(start+1, n, k, comb, result)
}
