package main

import "fmt"

func main() {
	elements := []int{0, 1, 2, 2, 4, 4, 4, 1}
	fmt.Println("frequentIndex: ", solutionTwo(elements))
}

func solutionOne(elements []int) int {
	frequentIndex := -1
	frequent := 0
	var counts [1000001]int

	for _, element := range elements {
		if element%2 == 0 {
			counts[element]++
		}
	}

	for element, count := range counts {
		if element%2 != 0 {
			continue
		}
		if count > frequent {
			frequent = count
			frequentIndex = element
		}
	}

	return frequentIndex
}

func solutionTwo(elements []int) int {
	frequentIndex := -1
	frequent := 0
	var counts [1000001]int

	for _, element := range elements {
		if element%2 == 0 {
			counts[element]++
			count := counts[element]
			if count > frequent {
				frequent = count
				frequentIndex = element
			}
		}
	}

	return frequentIndex
}
