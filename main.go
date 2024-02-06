package main

import "fmt"

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	maxLeft := height[0]
	maxRight := height[len(height)-1]
	water := 0

	for left < right {
		maxLeft = Max(maxLeft, height[left])
		maxRight = Max(maxRight, height[right])
		if maxLeft <= maxRight {
			water += maxLeft - height[left]
			left++
		} else {
			water += maxRight - height[right]
			right--
		}
	}
	return water
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))

}
