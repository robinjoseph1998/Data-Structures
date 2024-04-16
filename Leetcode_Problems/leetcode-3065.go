package main

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	minOp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			minOp++
		}
	}
	return minOp
}

func main() {

	nums := []int{2, 11, 10, 1, 3}
	k := 10
	fmt.Println(minOperations(nums, k))
}
