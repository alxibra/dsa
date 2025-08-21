package main

import (
	"fmt"
)

func main() {
	nums := []int{-4, -1, -1, 0, 1, 2}
	target := 0
	// result := [][]int{}
	// dsahelper.PrintSlice(nums)

	// for i := range nums[:len(nums)-2] {
	// 	if i > 0 && nums[i] == nums[i-1] {
	// 		continue
	// 	}
	// 	left, right := i+1, len(nums)-1
	//
	// 	for left < right {
	// 		sum := nums[i] + nums[left] + nums[right]
	// 		if sum == target {
	// 			result = append(result, []int{nums[i], nums[left], nums[right]})
	// 			left++
	// 			right--
	// 			continue
	// 		}
	//
	// 		if sum < target {
	// 			left++
	// 			continue
	// 		}
	//
	// 		right--
	// 	}
	// }

	fmt.Println("result: ", solutionOne(target, nums))
}

func solutionOne(target int, nums []int) [][]int {
	var result [][]int
	// dsahelper.PrintSlice(nums)
	fmt.Println("nums: ", nums)
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
			fmt.Printf("first + second + third = sum: %d+%d+%d = %d\n", first, second, third, sum)
			if sum == target {
				result = append(result, []int{first, second, third})
				left++
				right--
			}

			if sum < target {
				left++
				continue
			}

			if sum > target {
				right--
			}
		}
		fmt.Println("")
	}
	return result
}
