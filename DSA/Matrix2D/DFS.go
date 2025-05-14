package main

// adjacent cooridnates
// x-1,y-1 x-1,y  x-1,y+1
// x,y-1 	x,y	   x,y+1
// x+1,y-1 x+1,y  x+1,y+1

// DX = [-1,0,1,0]
// DY = [0,1,0,-1]

func DFSFloodFill(mat2D [][]int, i, j, existingColor, newColor int) {
	// Base case
	if !isValidCell(mat2D, i, j, existingColor, newColor) {
		return
	}
	mat2D[i][j] = newColor
	DFSFloodFill(mat2D, i-1, j, existingColor, newColor)
	DFSFloodFill(mat2D, i, j+1, existingColor, newColor)
	DFSFloodFill(mat2D, i+1, j, existingColor, newColor)
	DFSFloodFill(mat2D, i, j-1, existingColor, newColor)
}

func isValidCell(mat2D [][]int, i, j int, existingColor, newColor int) bool {
	if i < 0 || j < 0 || i >= len(mat2D) || j >= len(mat2D[0]) || mat2D[i][j] != existingColor {
		return false
	}
	return true
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	existingColor := image[sr][sc]
	if image[sr][sc] != newColor {
		DFSFloodFill(image, sr, sc, existingColor, newColor)
	}
	return image
}
