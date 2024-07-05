package main

import (
	"fmt"
)

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			if nums[i] == nums[i+1] {
				nums[i] = nums[i] * 2
				nums[i+1] = 0

			}
		}
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func main() {
	nums := []int{1, 2, 2, 1, 1, 0}

	fmt.Println(applyOperations(nums))
}
