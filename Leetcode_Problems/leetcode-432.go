package main

import (
	"fmt"
	"sort"
)

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	Occ := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] == k && i != j {
				Occ[nums[i]+nums[j]] = true
			}
		}
	}
	return len(Occ)
}

func main() {
	nums := []int{-1, -2, -3}
	k := 1
	fmt.Println(findPairs(nums, k))
}
