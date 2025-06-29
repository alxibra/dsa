package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 50, 99, 100}
	fmt.Println("missingNumbers: ", solutionOne(nums))

}

func solutionOne(nums []int) []int {
	var counts [101]int
	for _, n := range nums {
		counts[n]++
	}

	var missingNumbers []int
	for index, count := range counts {
		if count == 0 {
			missingNumbers = append(missingNumbers, index)
		}
	}
	return missingNumbers
}
