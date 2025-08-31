package main

import (
	"fmt"
)

// Input:
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
//
// Output:
// nums1 = [1,2,2,3,5,6]

func main() {
	// nums1 := []int{1, 2, 3, 0, 0, 0}
	// nums2 := []int{2, 5, 6}
	// n1 := 3
	// n2 := 3
	nums1 := []int{4, 7, 9, 0, 0, 0, 0}
	nums2 := []int{1, 2, 5, 8}
	n1 := 3
	n2 := 4
	fmt.Println("nums1: ", solutionOne(n1, n2, nums1, nums2))
}

func solutionOne(n1, n2 int, nums1, nums2 []int) []int {
	p1 := n1 - 1
	p2 := n2 - 1

	insertP := n1 + n2 - 1

	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[insertP] = nums1[p1]
			p1--
		} else {
			nums1[insertP] = nums2[p2]
			p2--
		}
		insertP--
	}
	for p2 >= 0 {
		nums1[insertP] = nums2[p2]
		insertP--
		p2--
	}

	return nums1
}
