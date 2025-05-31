package main

func getMaximumGold(grid [][]int) int {
	finalSum := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > 0 {
				findMaxGoldPath(grid, i, j, grid[i][j], &finalSum)
			}
		}
	}
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
		// fmt.Println(*finalSum)
	}

	// blocking the cell/ marking it as visited
	tmp := grid[i][j]
	grid[i][j] = 0

	findMaxGoldPath(grid, i+1, j, result+tmp, finalSum)
	findMaxGoldPath(grid, i-1, j, result+tmp, finalSum)
	findMaxGoldPath(grid, i, j+1, result+tmp, finalSum)
	findMaxGoldPath(grid, i, j-1, result+tmp, finalSum)

	// mark unvisted
	grid[i][j] = tmp
}
