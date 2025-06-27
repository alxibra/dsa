package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "The quick brown fox jumps over the lazy dog"
	if solutionOne(s) {
		fmt.Println("panagram")
		return
	}
}

func solutionOne(s string) bool {
	// Return early if the string is empty or
	// too short to be a pangram (needs at least 26 letters)
	if s == "" && len(s) < 26 {
		return false
	}

	s = strings.ToLower(s)
	var count [26]int

	for _, v := range s {
		// Skip non-alphabetic characters.
		if v < 'a' || v > 'z' {
			continue
		}
		count[v-'a']++
	}

	for _, v := range count {
		// If any character count is 0, it's not a pangram.
		if v == 0 {
			return false
		}
	}
	return true
}
