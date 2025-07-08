package main

func main() {
	nums := []int{-2, -1, 2, 1}
	k := 1

	prefixSum := 0
	maxLen := 0
	prefixMap := make(map[int]int) // map: prefixSum → index

	for i, num := range nums {
		prefixSum += num
		if prefixSum == k {
			maxLen = i + 1
		}
		if idx, found := prefixMap[prefixSum-k]; found {
			maxLen = max(maxLen, i-idx)
		}
		if _, exists := prefixMap[prefixSum]; !exists {
			prefixMap[prefixSum] = i
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
