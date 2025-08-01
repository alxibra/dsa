package main

import "fmt"

func main() {
	arr := []int{0, -1, 1, 100}
	fmt.Printf("solutionOne: %d\n", solutionOne(arr))
}

func solutionOne(arr []int) int {
	// If the array has less than 3 elements, it cannot be split into three non-empty parts.
	if len(arr) < 3 {
		return 0
	}

	// Calculate the total sum of the array.
	total := 0
	for _, e := range arr {
		total += e
	}

	// If the total sum is not divisible by 3,
	// it's impossible to split the array into three parts with equal sum.
	if total%3 != 0 {
		return 0
	}

	// We want to split the array into 3 parts with the same sum.
	// So each part must have sum = total / 3
	target := total / 3

	runningSum, firstPartCount, ways := 0, 0, 0

	for i := range arr[:len(arr)-1] {
		runningSum += arr[i]

		// If we've reached 2×target, it means the first two parts sum to 2×target,
		// so the third part must be target too.
		// Add all the ways we saw a prefix equal to target before this.
		if runningSum == 2*target {
			ways += firstPartCount
		}

		// Count the number of times we see a prefix equal to target
		if runningSum == target {
			firstPartCount++
		}
	}

	return ways
}
