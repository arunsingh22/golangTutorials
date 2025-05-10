package main

// traverse only the lower triangular matrix
// this works only with NxN matrix
func InPlaceTranspose(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < i; j++ {
			// swap the elements
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}
	print2DArray(mat)
}

// works with any NxM matrix
func Transpose(mat [][]int) {
	row_len := len(mat)
	col_len := len(mat[0])

	result := make([][]int, col_len) // row_len becomes col_len for a Transposed matrix

	for idx := range result {
		result[idx] = make([]int, row_len)
	}

	for i := 0; i < row_len; i++ {
		for j := 0; j < len(mat[i]); j++ {
			result[j][i] = mat[i][j]
		}
	}
	print2DArray(result)
}
