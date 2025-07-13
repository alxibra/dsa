package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("nums: ", nums)
	// fmt.Printf("pivotIndex: %d\n", solutionOne(nums))
	// fmt.Printf("pivotIndex: %d\n", solutionTwo(nums))
	fmt.Printf("pivotIndex: %d\n", solutionThree(nums))
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

func solutionTwo(nums []int) int {
	pivotIndex := -1
	prefixMap := make(map[int]int)

	// Improvement: Reuse prefixMap variable to avoid extra space for prefix sum.
	for index, num := range nums {
		if index == 0 {
			prefixMap[index] = num
			continue
		}
		prefixMap[index] = prefixMap[index-1] + num
	}

	leftSum, rightSum := 0, 0

	for key, value := range prefixMap {
		if key == 0 {
			continue
		}
		leftSum = value
		rightSum = prefixMap[len(prefixMap)-1] - prefixMap[key-1]
		if leftSum == rightSum {
			return key
		}
	}
	return pivotIndex
}

func solutionThree(nums []int) int {
	pivotIndex := -1
	prefixMap := make([]int, len(nums))

	for index, num := range nums {
		if index == 0 {
			prefixMap[index] = num
			continue
		}
		prefixMap[index] = prefixMap[index-1] + num
	}

	var leftSum int
	var rightSum int
	for index, num := range prefixMap {
		if index == 0 {
			continue
		}
		leftSum = num
		rightSum = prefixMap[len(prefixMap)-1] - prefixMap[index-1]
		if leftSum == rightSum {
			return index
		}
	}
	return pivotIndex
}
