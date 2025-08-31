package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println("max: ", solutionOne(nums))
}

func solutionOne(nums []int) int {

	max, left, right := 0, 0, len(nums)-1

	for left < right {

		currentArea := getArea(left, right, nums)

		if currentArea > max {
			max = currentArea
		}

		if nums[left] < nums[right] {
			left++
			continue
		}

		right--
	}
	return max
}

func getArea(left, right int, nums []int) int {
	minHeight := getMinimum(left, right, nums)
	width := right - left
	return minHeight * width
}

func getMinimum(left, right int, nums []int) int {
	if nums[left] < nums[right] {
		return nums[left]
	}

	return nums[right]
}
