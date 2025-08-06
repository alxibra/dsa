package main

import (
	"fmt"
	"math"
)

func main() {
	target := 7
	arr := []int{2, 3, 1, 2, 4, 3}
	fmt.Printf("minimum length sub array with target: %d\n", solutionOne(arr, target))
}

func solutionOne(arr []int, target int) int {
	prefixSum, length := 0, math.MaxInt
	hashTable := make(map[int]int, len(arr))

	for index, v := range arr {
		prefixSum += v

		if e, ok := hashTable[prefixSum-target]; ok {
			if length > (index - e) {
				length = (index - e)
			}
		}
		hashTable[prefixSum] = index
	}
	if length == math.MaxInt {
		length = 0
	}
	return length
}
