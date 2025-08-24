package main

import "fmt"

func main() {
	height := []int{4, 2, 0, 3, 2, 5}
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	trap := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				trap += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				trap += rightMax - height[right]
			}
			right--
		}
	}

	fmt.Println("trap: ", trap)
}
