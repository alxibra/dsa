package main

import "fmt"

func main() {
	s := "swiss"

	fmt.Printf("the first character is: %s\n", solutionOne(s))

}

func solutionOne(s string) string {
	// a-z is 26 character
	var count [26]int

	// count each character in string s
	for _, v := range s {
		count[v-'a']++
	}

	// Find the first non-repeating character.
	for _, v := range s {
		if count[v-'a'] == 1 {
			return string(v)
		}
	}

	return "-"
}
