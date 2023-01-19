package main

import "fmt"

func Solution(S string) string {
	// write your code in Go 1.13.8
	i := 0
	j := len(S) - 1

	res := make([]byte, len(S))
	for i < j {
		if S[i] == '?' && S[j] != '?' {
			res[i] = S[j]
			i++
			j--
			fmt.Println(string(res))
		} else if S[i] != '?' && S[j] == '?' {
			res[j] = S[i]
			i++
			j--
			fmt.Println(string(res))
		} else if S[i] == '?' && S[j] == '?' { 
			res[i] = 'a'
			res[j] = 'a'
			i++
			j--
			fmt.Println(string(res))
		} else if S[i] != S[j] {
			return "NO"
		} else {
			// both are same
			res[i] = S[i]
			res[j] = S[j]
			i++
			j--
			fmt.Println(string(res))
		}
	}
	// fmt.Println(res.String())
	fmt.Println(string(res))
	return string(res)
}

func main() {
	Solution("?ab??a")
}
