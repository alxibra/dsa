package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, -1, -1, 0, 1, 2}
	target := 0
	fmt.Println("result: ", solutionOne(target, nums))
}

func solutionOne(target int, nums []int) [][]int {
	var result [][]int
	for i := range nums[:len(nums)-2] {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		first := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			second := nums[left]
			third := nums[right]
			sum := first + second + third
			if sum == target {
				result = append(result, []int{first, second, third})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
			if sum < target {
				left++
			}
			if sum > target {
				right--
			}
		}
	}
	return result
}
