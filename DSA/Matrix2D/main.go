package main

import "fmt"

func main() {

	// 🔹 Memory Layout and Storage in Go
	// Go stores slices indirectly:
	// A slice is a descriptor containing a pointer to an underlying array, a length, and a capacity.
	// When you have a 3D slice like [][][]int, you have a slice of slice of slices — each level may point to a different underlying array.
	// So it is not stored as a contiguous block (like in C/C++ fixed-size multi-dimensional arrays). Instead:
	// mat3D points to a slice of [][]int
	// Each [][]int points to a slice of []int
	// Each []int points to actual integer values

	// NOTE: This makes 3D slices flexible but less cache-friendly and
	// potentially fragmented in memory.

	// declaring a 2D matrix
	// mat2D := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}
	// findDiagonalOrder(mat2D)
	// InPlaceTranspose(mat2D)
	// Transpose(mat2D)
	// print2DArray(mat2D)

	// mat3D := [][][]int{{{1, 2}, {3, 4}, {5, 6}}}
	// Transpose(mat3D)
	// print3DArray(mat3D)

	mat := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	floodFill(mat, 1, 1, 2)
	fmt.Println(mat)

}

func print3DArray(mat [][][]int) {
	row_len := len(mat)
	col_len := len(mat[0])
	high_len := len(mat[0][0])

	fmt.Println("3D array dim: ", row_len, col_len, high_len)
}

// NOTE: unlike C mat doesn't get reduced to a pointer
func print2DArray(mat [][]int) {
	row_len := len(mat)
	col_len := len(mat[0])
	fmt.Println(row_len, col_len)

	// standard way: accessing element by element
	// fmt.Println("Standard Way:")
	// for i := 0; i < row_len; i++ {
	// 	for j := 0; j < col_len; j++ {
	// 		fmt.Print(mat[i][j], " ")
	// 	}
	// 	fmt.Println("")
	// }

	// golang way: accessing row by row
	fmt.Println("Accessing row wise:")
	for _, row := range mat {
		// fmt.Println(idx, row)
		for _, val := range row {
			fmt.Print(val, ",")
		}
		fmt.Println()
	}
}
