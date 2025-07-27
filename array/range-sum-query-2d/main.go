package main

import "fmt"

type NumMatrix struct {
	prefix [][]int
}

func main() {
	row1, column1, row2, column2 := 2, 1, 4, 3
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	printMatrix(matrix)
	nm := constructor(matrix)
	fmt.Println("")
	fmt.Println("\tPrefix Sum Region")
	printMatrix(nm.prefix)
	fmt.Println("")
	fmt.Printf("sumRegion: %d\n", nm.sumRegion(row1, column1, row2, column2))
}

func (nm NumMatrix) sumRegion(row1, column1, row2, column2 int) int {
	fmt.Printf("row_1: %d, column_1: %d, row_2: %d, column_2: %d\n", row1, column1, row2, column2)
	p := nm.prefix
	return p[row2+1][column2+1] -
		p[row1][column2+1] -
		p[row2+1][column1] +
		p[row1][column1]
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
	// Column headers
	fmt.Printf("     ")
	for j := range matrix[0] {
		fmt.Printf(" %2d ", j)
	}
	fmt.Println()

	fmt.Printf("     ")
	for range matrix[0] {
		fmt.Print("----")
	}
	fmt.Println()

	// Row data with range
	for i, row := range matrix {
		fmt.Printf(" %2d |", i)
		for _, val := range row {
			fmt.Printf(" %2d ", val)
		}
		fmt.Println()
	}
}
