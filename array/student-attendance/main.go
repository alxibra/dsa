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
	var late int

	for _, character := range records {
		cm := characterMap(character)
		counts[cm]++

		// Check for more than one absence (A)
		if cm == 1 && counts[cm] > 1 {
			return false
		}

		// Reset late counter when encountering Present (P) or Absent (A)
		if cm == 0 || cm == 1 {
			late = 0
		}

		// Track consecutive late days (L)
		if cm == 2 {
			late++
			// Fail if 3 consecutive late days
			if late == 3 {
				return false
			}
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
