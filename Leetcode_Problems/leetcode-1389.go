package main

import (
	"fmt"
	"slices"
)

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		target = slices.Insert(target, index[i], nums[i])
	}
	return target
}

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	fmt.Println(createTargetArray(nums, index))
}
