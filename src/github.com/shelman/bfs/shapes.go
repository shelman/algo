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

func largestShape(grid [][]int8) shape {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return shape{}
	}

	rows := len(grid)
	cols := len(grid[0])

	largestShape := shape{}
	visited := map[coordinate]bool{}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			startPoint := coordinate{row, col}
			if grid[row][col] == 0 || visited[startPoint] {
				continue
			}

			queue := []coordinate{startPoint}
			currentShape := shape{}
			for len(queue) > 0 {
				point := queue[0]
				queue = queue[1:]

				if grid[point.row][point.col] == 1 && !visited[point] {
					currentShape = append(currentShape, point)
					visited[point] = true

					for _, neighbor := range []coordinate{
						{point.row-1, point.col},
						{point.row+1, point.col},
						{point.row, point.col-1},
						{point.row, point.col+1},
					} {
						if !visited[neighbor] && neighbor.row >= 0 && neighbor.row < rows &&
							neighbor.col >= 0 && neighbor.col < cols &&
							grid[neighbor.row][neighbor.col] == 1 {
								queue = append(queue, neighbor)
						}
					}
				}
			}

			if len(currentShape) > len(largestShape) {
				largestShape = currentShape
			}

		}
	}

	return largestShape
}