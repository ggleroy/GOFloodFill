package main

import "fmt"

type CellCoordinates struct {
	X, Y int
}

func floodFillIterative(matrix [][]int, x, y, regionNumber int) {

	queue := []CellCoordinates{{X: x, Y: y}}
	matrix[x][y] = regionNumber

	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		directions := []CellCoordinates{
			{X: currentPoint.X - 1, Y: currentPoint.Y},
			{X: currentPoint.X + 1, Y: currentPoint.Y},
			{X: currentPoint.X, Y: currentPoint.Y - 1},
			{X: currentPoint.X, Y: currentPoint.Y + 1},
		}

		for _, dir := range directions {
			nx, ny := dir.X, dir.Y

			if nx >= 0 && nx < len(matrix) && ny >= 0 && ny < len(matrix[0]) && matrix[nx][ny] == 0 {
				matrix[nx][ny] = regionNumber
				queue = append(queue, CellCoordinates{X: nx, Y: ny})
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 1, 0, 0, 1},
		{0, 1, 1, 0, 1},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 0},
		{0, 0, 1, 0, 0},
	}

	fmt.Println("Initial Matrix:")
	printMatrix(matrix)

	color := 2
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				floodFillIterative(matrix, i, j, color)
				color++
			}
		}
	}

	fmt.Println("\nFinal Matrix after Flood Fill:")
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}