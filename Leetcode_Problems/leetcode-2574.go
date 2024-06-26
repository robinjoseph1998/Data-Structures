package main

import "fmt"

func leftRightDifference(nums []int) []int {
	right := []int{}
	left := []int{0}
	leftAdder := 0
	for i := 0; i < len(nums)-1; i++ {
		leftAdder += nums[i]
		left = append(left, leftAdder)
	}
	rightAdder := 0
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {
			rightAdder += nums[j]
		}
		right = append(right, rightAdder)
		rightAdder = 0
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		res = append(res, Helper(right[i], left[i]))
	}
	return res
}

func Helper(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	nums := []int{10, 4, 8, 3}
	fmt.Println(leftRightDifference(nums))
}
