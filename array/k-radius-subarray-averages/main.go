package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	fmt.Println("solutionOne: ", solutionOne(k, arr))
}

func solutionOne(k int, arr []int) []int {
	prefixSum := 0
	prefixSumTable := make([]int, len(arr))
	solution := make([]int, len(arr))

	if k == 0 {
		return arr
	}

	window := (2 * k) + 1
	if window > len(arr) {
		return arr
	}
	for i, v := range arr {
		prefixSum += v
		prefixSumTable[i] = prefixSum
	}

	for i := range solution {
		solution[i] = -1
	}

	for i := k; i < len(arr)-k; i++ {
		if i == k {
			average := prefixSumTable[i+k] / ((2 * k) + 1)
			solution[i] = average
			continue
		}
		sum := prefixSumTable[i+k] - prefixSumTable[i-k-1]
		average := (sum) / ((2 * k) + 1)
		solution[i] = average
	}
	return solution
}

func printSlice(nums []int) {
	if len(nums) == 0 {
		fmt.Println("(empty slice)")
		return
	}

	// 1. Find the widest cell (max of index‐width & value‐width).
	cell := 0
	for i, v := range nums {
		if w := len(strconv.Itoa(i)); w > cell {
			cell = w
		}
		if w := len(strconv.Itoa(v)); w > cell {
			cell = w
		}
	}
	cell += 1

	printRow := func(label string, values []int) {
		fmt.Printf("%-5s |", label)
		for _, x := range values {
			fmt.Printf("%*d", cell, x)
		}
		fmt.Println()
	}

	// 3. Index row.
	indices := make([]int, len(nums))
	for i := range nums {
		indices[i] = i
	}
	printRow("index", indices)

	// 4. Separator (---+------).
	fmt.Print(strings.Repeat("-", 5), "-+")
	fmt.Println(strings.Repeat("-", cell*len(nums)))

	// 5. Value row.
	printRow("value", nums)
}
