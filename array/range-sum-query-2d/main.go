package main

import "fmt"

func main() {
	// row1, column1, row2, column2 := 2, 1, 4, 3
	// matrix := [][]int{
	// 	{3, 0, 1, 4, 2},
	// 	{5, 6, 3, 2, 1},
	// 	{1, 2, 0, 1, 5},
	// 	{4, 1, 0, 1, 7},
	// 	{1, 0, 3, 0, 5},
	// }
	row1, column1, row2, column2 := 0, 0, 1, 1
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println("solutionOne: ", solutionOne(row1, column1, row2, column2, matrix))
}

func solutionOne(r1, c1, r2, c2 int, m [][]int) int {
	printMatrix(m)
	ps := generatePrefixSum2D(m)
	fmt.Println("")
	printMatrix(ps)
	return sumOfRegion(r1, c1, r2, c2, ps)
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

func generatePrefixSum2D(m [][]int) [][]int {
	rowCount := len(m)
	columnCount := len(m[0])

	prefixSum := make([][]int, rowCount+1)

	for index := range prefixSum {
		prefixSum[index] = make([]int, columnCount+1)
	}

	for i := range rowCount {
		for j := range columnCount {
			prefixSum[i+1][j+1] = m[i][j] +
				prefixSum[i][j+1] +
				prefixSum[i+1][j] -
				prefixSum[i][j]
		}
	}

	return prefixSum
}

func sumOfRegion(r1, c1, r2, c2 int, ps [][]int) int {
	return ps[r2+1][c2+1] -
		ps[r1][c2+1] -
		ps[r2+1][c1] +
		ps[r1][c1]
}
