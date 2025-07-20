package main

import "fmt"

func main() {
	// nums := []int{0, 1, 1}
	nums := []int{0, 1, 2}
	// fmt.Printf("countEven: %d\n", solutionCountEvenNum(nums))
	fmt.Printf("SolutionCountOddNum: %d\n", SolutionCountOddNum(nums))
}

func solutionCountEvenNum(nums []int) int {
	var count, prefixSum, mod = 0, 0, 0
	hashTable := map[int]int{0: 1}
	for _, num := range nums {
		prefixSum += num
		mod = prefixSum % 2
		hashTable[mod]++
	}
	return count
}

func SolutionCountOddNum(nums []int) int {
	var count, prefixSum, mod = 0, 0, 0

	hashTable := make(map[int]int)

	hashTable[0] = 1

	for _, num := range nums {

		prefixSum += num

		mod = prefixSum % 2

		count += hashTable[mod-1]

		hashTable[mod]++

	}
	return count
}
