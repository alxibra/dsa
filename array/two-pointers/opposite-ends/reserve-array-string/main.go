package main

import "fmt"

func main() {
	str := "hello"

	fmt.Println("solution: ", solutionOne(str))
}

func solutionOne(str string) string {
	arr := []byte(str)

	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	return string(arr)
}
