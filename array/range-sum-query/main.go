package main

import (
	"fmt"
)

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	// fmt.Println("sum: ", solutionOne(2, 4, nums))
	fmt.Println("sum: ", solutionTwo(2, 4, nums))
	// fmt.Println("sum: ", solutionThree(2, 4, nums))
}

// NOTE: use this solution for better
func solutionOne(left, right int, nums []int) int {
	// Create a prefix sum array with an extra slot at the beginning
	// prefix[0] is set to 0 to handle sum from index 0 gracefully
	// This allows us to use the formula: prefix[r+1] - prefix[l]
	var prefix = make([]int, len(nums)+1)

	for i, num := range nums {
		// prefix[i+1] = sum of nums[0..i]
		prefix[i+1] = prefix[i] + num
	}

	// Return the sum of elements from nums[left] to nums[right] inclusive
	return prefix[right+1] - prefix[left]
}

func solutionTwo(left, right int, nums []int) int {
	for i, num := range nums {
		if i == 0 {
			continue
		}
		nums[i] = nums[i-1] + num
	}
	fmt.Printf("nums: %v\n", nums)
	if left == 0 {
		return nums[right]
	}
	fmt.Printf("right: %d\n", nums[right])
	fmt.Printf("left: %d\n", nums[left-1])
	return nums[right] - nums[left-1]
}

func solutionThree(left, right int, nums []int) int {

	prefixSum := make(map[int]int)

	for index, num := range nums {

		if index == 0 {
			prefixSum[index] = num
			continue
		}

		prefixSum[index] = prefixSum[index-1] + num

	}

	if left == 0 {
		return prefixSum[right]
	}

	return prefixSum[right] - prefixSum[left-1]
}
