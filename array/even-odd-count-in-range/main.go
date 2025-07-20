package main

import "fmt"

func main() {
	nums := []int{0, 1, 1}
	fmt.Printf("countEven: %d\n", solutionCountEvenNum(nums))
}

func solutionCountEvenNum(nums []int) int {
	var count, prefixSum, mod = 0, 0, 0
	hashTable := map[int]int{0: 1}
	for _, num := range nums {
		prefixSum += num
		mod = prefixSum % 2
		count += hashTable[mod]

		hashTable[mod]++
	}
	return count
}
