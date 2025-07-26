package main

import "fmt"

type NumMatrix struct {
	prefix [][]int
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	printMatrixWithIndexes(matrix)
	nm := constructor(matrix)
	fmt.Println("")
	printMatrixWithIndexes(nm.prefix)
}

func constructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	column := len(matrix[0])

	prefixSum := make([][]int, row+1)

	for index := range prefixSum {
		prefixSum[index] = make([]int, column+1)
	}

	for i := range row {
		for j := range column {
			prefixSum[i+1][j+1] =
				matrix[i][j] +
					prefixSum[i][j+1] +
					prefixSum[i+1][j] -
					prefixSum[i][j]
		}
	}

	return NumMatrix{prefixSum}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Print("[ ")
		for _, val := range row {
			fmt.Printf("%3d ", val) // align with 3-digit width
		}
		fmt.Println("]")
	}
}

func printMatrixWithIndexes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	// Column headers
	fmt.Printf("     ")
	for j := 0; j < cols; j++ {
		fmt.Printf(" %2d ", j)
	}
	fmt.Println()

	fmt.Printf("     ")
	for j := 0; j < cols; j++ {
		fmt.Print("----")
	}
	fmt.Println()

	// Row data
	for i := 0; i < rows; i++ {
		fmt.Printf(" %2d |", i)
		for j := 0; j < cols; j++ {
			fmt.Printf(" %2d ", matrix[i][j])
		}
		fmt.Println()
	}
}
