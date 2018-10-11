package main

import "fmt"

func main() {
	grid := [][]int8{
		{1, 0, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 0, 0},
	}
	fmt.Printf("#%v\n", largestShape(grid))
}

type coordinate struct {
	row int
	col int
}

type shape []coordinate

func buildShape(outputShape *shape, point coordinate, visited map[coordinate]bool, grid [][]int8,
	rows int, cols int) {
	if visited[point] || point.row < 0 || point.row >= rows ||
		point.col < 0 || point.col >= cols || grid[point.row][point.col] == 0 {
		return
	}

	*outputShape = append(*outputShape, point)
	visited[point] = true

	buildShape(outputShape, coordinate{point.row-1, point.col}, visited, grid, rows, cols)
	buildShape(outputShape, coordinate{point.row+1, point.col}, visited, grid, rows, cols)
	buildShape(outputShape, coordinate{point.row, point.col-1}, visited, grid, rows, cols)
	buildShape(outputShape, coordinate{point.row, point.col+1}, visited, grid, rows, cols)
}

func largestShape(grid [][]int8) shape {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return shape{}
	}

	rows := len(grid)
	cols := len(grid[0])

	largestShape := shape{}
	visited := map[coordinate]bool{}

	for row := 0; row < rows; row++  {
		for col := 0; col < cols; col++ {
			startCoord := coordinate{row, col}
			currentShape := shape{}
			buildShape(&currentShape, startCoord, visited, grid, rows, cols)

			if len(currentShape) > len(largestShape) {
				largestShape = currentShape
			}
		}
	}

	return largestShape
}