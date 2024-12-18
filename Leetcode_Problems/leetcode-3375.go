package main

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	if nums[0] < k {
		return -1
	}
	operations := 0
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i] != nums[i-1] {
			operations++
		}
	}
	if nums[0] > k {
		operations++
	}
	return operations
}

func main() {
	nums := []int{5, 2, 5, 4, 5}
	k := 2
	fmt.Println(minOperations(nums, k))
}
