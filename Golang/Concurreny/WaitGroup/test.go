package main

import "fmt"

func main() {

	// while ranging if only 1 variable is used it will be idx by default.
	dict := map[int]int{}

	for i := 0; i < 10; i++ {
		dict[i] = i + 1
	}

	for k, v := range dict {
		fmt.Println(k, v)
	}
	fmt.Println("--------------------")
	for x := range dict {
		fmt.Println(x)
	}

	// ans := []int{5, 6, 7, 8}
	// for x := range ans {
	// 	fmt.Println(x)
	// }
}
