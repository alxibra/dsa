package main

import (
	"fmt"
)

func main() {
	// records := "PPALLP"
	records := "PPPALLL"
	if solutionOne(records) {
		fmt.Println("Award")
		return
	}
	fmt.Println("No Award")
}

func solutionOne(records string) bool {
	var counts [3]int
	var consecutive int

	for _, character := range records {
		cm := characterMap(character)
		counts[cm]++
		if cm == 1 && counts[cm] > 1 {
			return false
		}
		if cm == 1 || cm == 0 {
			consecutive = 0
		}
		if cm == 2 {
			consecutive++
		}
		if consecutive == 3 {
			return false
		}
	}
	return true
}

func characterMap(c rune) int {
	switch c {
	case 'P':
		return 0
	case 'A':
		return 1
	case 'L':
		return 2
	default:
		return 100
	}
}
