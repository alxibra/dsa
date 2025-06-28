package main

import (
	"fmt"
)

func main() {
	birds := []int{1, 4, 4, 4, 5, 3, 3}
	maxID := solutionTwo(birds)
	fmt.Println("max id: ", maxID)
}

// solutionOne use 2 step approach
// separation of concern, easy to reason about, maintain and debug
func solutionOne(birds []int) int {
	var counts [6]int
	var max int
	var maxID int

	for _, ID := range birds {
		counts[ID]++
	}

	for ID, count := range counts {
		if count > max {
			max = count
			maxID = ID
		}
	}
	return maxID
}

// solutionOne use 1 step approach
// only one looping more efficient
func solutionTwo(birds []int) int {
	var counts [6]int
	var max int
	var maxID int

	for _, ID := range birds {
		counts[ID]++
		if counts[ID] > max {
			max = counts[ID]
			maxID = ID
		}
	}
	return maxID
}
