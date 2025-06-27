package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "silent"
	l := "listene"

	if solutionThree(s, l) {
		fmt.Println("anagram")
		return
	}
	fmt.Println("not anagram")
}

func solutionOne(s, l string) bool {
	if len(s) != len(l) {
		return false
	}
	// trivia
	// strings in Go are essentially slices of bytes ([]byte)
	// byte is an alias for uint8, representing a single byte (0 to 255).
	// ASCII characters are represented by byte values 0 to 127.

	// why 128? Assuming extended ASCII (0-127) or basic ASCII.
	var count [128]int

	// Loop through each character in the strings
	// Since both strings are the same length, we can use a single loop
	// For each character in string `s`, increment its count in the `count` array
	// For each corresponding character in string `l`, decrement its count
	// If `s` and `l` are anagrams, the increments and decrements will cancel out,
	// resulting in all elements of the `count` array being zero at the end.
	for i := range s {
		count[s[i]]++ // count up character from s
		count[l[i]]-- // count down character from l
	}

	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func solutionTwo(s, l string) bool {
	if len(s) != len(l) {
		return false
	}

	// why 26, because a-z is 26 count
	// we assume all small case
	// compare to solutionOne,
	// we can improve memory usage by using an array of size 26 instead of 128.
	// 96 bytes will be save
	var count [26]int

	// Why subtract 'a'?
	// This converts the character (assuming lowercase ASCII letters) to an index from 0-25.
	// 'a' has an ASCII value of 97.  Subtracting 'a' maps 'a' to index 0, 'b' to index 1, and so on.
	for i := range s {
		count[s[i]-'a']++
		count[l[i]-'a']--
	}
	fmt.Println('a')

	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}

func solutionThree(s, l string) bool {
	if len(s) != len(l) {
		return false
	}
	// Before sorting, the string must be split into a slice of runes.
	as := strings.Split(s, "")
	al := strings.Split(l, "")

	sort.Strings(as)
	sort.Strings(al)

	// After splitting and sorting, iterate through s and l. If any characters mismatch,
	// they are not anagrams.
	for i := range as {
		if as[i] != al[i] {
			return false
		}
	}
	return true
}
