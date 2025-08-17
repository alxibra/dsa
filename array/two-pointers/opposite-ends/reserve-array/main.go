package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("nums: \t\t", nums)
	fmt.Println("reverse: \t", solutionOne(nums))
}

func solutionOne(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
