package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("nums: ", nums)
	fmt.Printf("pivotIndex: %d\n", solutionOne(nums))
}

func solutionOne(nums []int) int {
	pivotIndex := -1
	var prefixSum int
	prefixMap := make(map[int]int)

	for index, num := range nums {
		prefixSum += num
		prefixMap[index] = prefixSum
	}

	for index, v := range prefixMap {
		leftSum := v
		rightIndex := len(prefixMap) - 1
		rightSum := prefixMap[rightIndex] - prefixMap[index-1]
		if rightSum == leftSum {
			pivotIndex = index
		}
	}
	return pivotIndex
}
