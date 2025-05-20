package main

import (
	"fmt"
)

const (
	D = "D"
	L = "L"
	R = "R"
	U = "U"
)

func main() {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1},
	}
	// maze := [][]int{
	// 	{1, 0},
	// 	{1, 1},
	// }
	paths := []string{}
	findAllPaths(maze, 0, 0, &paths, "")
	fmt.Println("Rat in Maze paths: ", paths)
}

func findAllPaths(maze [][]int, i, j int, paths *[]string, path string) {
	// BC 1
	m := len(maze)
	n := len(maze[0])
	if i < 0 || j < 0 || i >= m || j >= n || maze[i][j] == 0 {
		return
	}
	// BC 2
	if i == m-1 && j == n-1 && maze[i][j] == 1 {
		*paths = append(*paths, path)
		return
	}
	// blocking the cell/ marking it as visited
	maze[i][j] = 0

	findAllPaths(maze, i+1, j, paths, path+D)
	findAllPaths(maze, i-1, j, paths, path+U)
	findAllPaths(maze, i, j+1, paths, path+R)
	findAllPaths(maze, i, j-1, paths, path+L)

	// mark unvisted
	maze[i][j] = 1
}
