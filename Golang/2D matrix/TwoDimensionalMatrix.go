package main

import "fmt"

func main() {

	// Declaring a 2D matrix
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(mat[i][j], " ")
		}
		fmt.Println("")
	}
	fmt.Println(mat)
}
