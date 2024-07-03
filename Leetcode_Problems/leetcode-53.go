package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	CurrentSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if CurrentSum < 0 {
			CurrentSum = nums[i]
		} else {
			CurrentSum += nums[i]
		}

		if CurrentSum > maxSum {
			maxSum = CurrentSum
		}
	}
	return maxSum
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
