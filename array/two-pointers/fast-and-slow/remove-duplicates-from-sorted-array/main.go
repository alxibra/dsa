package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3, 4, 5, 5}

	fmt.Println("solution: ", solutionOne(nums))
}

func solutionOne(nums []int) []int {
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return nums[:slow+1]
}
