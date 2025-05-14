package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				if DFS(board, i, j, 0, word) {
					return true
				}
			}
		}
	}
	return false
}

func DFS(board [][]byte, i, j int, k int, word string) bool {
	if k == len(word) {
		return true
	}
	if !isValidMove(board, i, j, k, word) {
		return false
	}

	// marking visited
	tmp := board[i][j]
	board[i][j] = ' '

	found := DFS(board, i+1, j, k+1, word) || DFS(board, i-1, j, k+1, word) ||
		DFS(board, i, j+1, k+1, word) || DFS(board, i, j-1, k+1, word)

	// backtrack the cell
	board[i][j] = tmp
	return found
}

func isValidMove(board [][]byte, i, j int, k int, word string) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] != word[k] {
		return false
	}
	return true
}
