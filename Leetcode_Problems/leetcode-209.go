package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	windowSum := 0
	minLength := math.MaxInt32
	start := 0

	for end := 0; end < n; end++ {
		windowSum += nums[end]

		for windowSum >= target {
			minLength = Min(minLength, end-start+1)
			windowSum -= nums[start]
			start++
		}

	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))
}
