package main

import "fmt"

// Input:  nums = [3,2,2,3], val = 3
// Output: 2
// Explanation: Your function should return length = 2,
// and the first two elements of nums should be 2 and 2.

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println("Output: ", solutionOne(val, nums))
}

func solutionOne(val int, nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return slow
}
