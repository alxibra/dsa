package main

import "fmt"

func main() {
	str := "racecar"

	fmt.Println("is it palindrom? ", solutionOne(str))
}

func solutionOne(str string) bool {
	arr := []byte(str)

	left, right := 0, len(arr)-1

	for left < right {
		if arr[left] != arr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
