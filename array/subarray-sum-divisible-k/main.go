package main

import "fmt"

func main() {
	// nums := []int{4, 5, 1}
	nums := []int{2, 3, 1, 6, 4}
	k := 5
	fmt.Printf("count: %d\n", solutionOne(nums, k))
}

func solutionOne(nums []int, k int) int {
	count, prefix, mod := 0, 0, 0
	modMap := make(map[int]int)
	// Handle prefixSum that is directly divisible by k (starting from index 0)
	modMap[0] = 1
	for _, num := range nums {
		prefix += num
		// Normalize the mod to be in [0, k-1]
		mod = ((prefix % k) + k) % k
		// Count how many times this mod has occurred before
		// If mod has been seen before, that means the subarray between
		// that previous index and now has a sum divisible by k
		count += modMap[mod]
		// Record this mod for future use
		modMap[mod]++
	}

	return count
}
