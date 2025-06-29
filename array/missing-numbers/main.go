package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 50, 99, 100}
	fmt.Println("missingNumbers: ", solutionTwo(nums))

}

func solutionOne(nums []int) []int {
	var counts [101]int
	for _, n := range nums {
		counts[n]++
	}

	var missingNumbers []int
	for index, count := range counts {
		// if count = 0, it must not present in nums
		if count == 0 {
			missingNumbers = append(missingNumbers, index)
		}
	}
	return missingNumbers
}

func solutionTwo(nums []int) []int {
	var seen [101]bool
	for _, n := range nums {
		seen[n] = true
	}

	var missingNumbers []int
	for index, found := range seen {
		// if not found, it must not present in nums
		if !found {
			missingNumbers = append(missingNumbers, index)
		}
	}
	return missingNumbers
}
