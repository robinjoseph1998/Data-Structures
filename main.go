package main

import (
	"fmt"
	"sort"
)

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			if nums[i] == nums[i+1] {
				return true
			}
		}
	}
	return false
}

func main() {

	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
