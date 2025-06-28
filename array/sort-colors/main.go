package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	nums = solutionOne(nums)
	fmt.Println(nums)
}

func solutionOne(nums []int) []int {
	var counts [3]int

	// Count occurrences of 0, 1, and 2
	for _, value := range nums {
		counts[value]++
	}

	fmt.Println(counts)
	index := 0
	// trivia
	// "in-place" means the algorithm does not use extra memory to store
	// a new copy of input

	// In-place overwrite nums with sorted colors based on the counts
	for color, count := range counts {
		for _ = range count {
			nums[index] = color
			index++
		}
	}
	return nums
}
