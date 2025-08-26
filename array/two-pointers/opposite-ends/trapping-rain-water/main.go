package main

import "fmt"

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	// left, right := 0, len(height)-1
	// leftMax, rightMax := 0, 0
	// trap := 0
	// for left < right {
	// 	if height[left] < height[right] {
	// 		if height[left] >= leftMax {
	// 			leftMax = height[left]
	// 		} else {
	// 			trap += leftMax - height[left]
	// 		}
	// 		left++
	// 	} else {
	// 		if height[right] > rightMax {
	// 			rightMax = height[right]
	// 		} else {
	// 			trap += rightMax - height[right]
	// 		}
	// 		right--
	// 	}
	// 	fmt.Println("")
	// }

	fmt.Println("trap: ", solutionOne(height))
}

func solutionOne(height []int) int {
	var trap int
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		// find the heigh of left and right
		leftHeight, rightHeight := height[left], height[right]
		// If leftHeight < rightHeight, we process the left side first
		// because the amount of water trapped is limited by the shorter bar.
		if leftHeight < rightHeight {
			if leftHeight > leftMax {
				// No water is trapped when leftHeight exceeds leftMax; update leftMax
				leftMax = leftHeight
			} else {
				// Water trapped is the difference between leftMax and leftHeight
				trap += leftMax - leftHeight
			}
			left++
		} else {
			if rightHeight > rightMax {
				// No water is trapped when rightHeight exceeds rightMax, update rightMax
				rightMax = rightHeight
			} else {
				// Water trapped is the difference between rightMax and leftHeight
				trap += rightMax - rightHeight
			}
			right--
		}
	}
	return trap
}
