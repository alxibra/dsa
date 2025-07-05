package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println("count: ", solutionOne(nums, k))
}

func solutionOne(nums []int, k int) int {
	var prefix, count int
	prefixMap := map[int]int{0: 1}
	for _, num := range nums {
		prefix += num
		// At this step, I want to know if there’s any earlier prefix sum that,
		// when removed from the current total, gives me exactly k.
		// If yes, I found a subarray (or more) ending here that sums to k.
		if val, found := prefixMap[prefix-k]; found {
			// I look up how many times prefix - k has occurred
			// and I add that many to my count.
			count += val
		}
		prefixMap[prefix] += 1
	}
	return count
}
