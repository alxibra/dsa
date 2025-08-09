package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	k := 3
	expected := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}
	results := make([]int, len(expected))
	prefixSumTable := make([]int, len(expected))

	fmt.Println("'arr'")
	printSlice(arr)
	fmt.Println("expected")
	printSlice(expected)
	fmt.Printf("k: %d\n", k)

	for i := range results {
		results[i] = -1
	}
	prefixSum := 0
	for i, v := range arr {
		prefixSum += v
		prefixSumTable[i] = prefixSum
	}

	for i := k; i < len(results)-k; i++ {
		var leftSum int
		left := i - k - 1
		leftSum = 0
		if left >= 0 {
			leftSum = prefixSumTable[left]
		}
		rightSum := prefixSumTable[i+k]
		results[i] = (rightSum - leftSum) / (2*k + 1)
	}

	fmt.Println("")
	fmt.Println("results")
	printSlice(results)
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
