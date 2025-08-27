package main

import "fmt"

// Given an integer array nums, move all the 0’s to the end of the array
// while maintaining the relative order of the non-zero elements.
// You must do this in-place without making a copy of the array.
// Input:  nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
func main() {
	nums := []int{0, 1, 0, 3, 12}

	fmt.Println("solution: ", solutionOne(nums))
}

func solutionOne(nums []int) []int {
	slow, fast := 0, 0

	for fast < len(nums) {

		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}

		fast++

	}

	return nums
}
