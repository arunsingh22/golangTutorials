package main

import "fmt"

func getMaximumGold(grid [][]int) int {
	result := -1
	finalSum := 0
	m := len(grid)
	n := len(grid[0])

	s, c := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 && grid[i][j] > int(result) {
				result = grid[i][j]
				s, c = i, j
			}
		}
	}
	fmt.Println(result)
	findMaxGoldPath(grid, s, c, result, &finalSum)
	return finalSum
}

func findMaxGoldPath(grid [][]int, i, j, result int, finalSum *int) {
	m := len(grid)
	n := len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] == 0 {
		return
	}
	if result+grid[i][j] > *finalSum {
		*finalSum = result + grid[i][j]
	}

	// blocking the cell/ marking it as visited
	grid[i][j] = 0

	findMaxGoldPath(grid, i+1, j, result+grid[i][j], finalSum)
	findMaxGoldPath(grid, i-1, j, result+grid[i][j], finalSum)
	findMaxGoldPath(grid, i, j+1, result+grid[i][j], finalSum)
	findMaxGoldPath(grid, i, j-1, result+grid[i][j], finalSum)

	// mark unvisted
	grid[i][j] = 1
}
