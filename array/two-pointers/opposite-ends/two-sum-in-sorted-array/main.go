package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println("result: ", solutionOne(nums, target))
}

func solutionOne(nums []int, target int) []int {
	var result []int
	left, right, sum := 0, len(nums)-1, 0

	if len(nums) < 2 {
		return result
	}

	for left < right {
		sum = nums[left] + nums[right]

		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			left++
		}

		if sum > target {
			right--
		}
	}
	return result
}
