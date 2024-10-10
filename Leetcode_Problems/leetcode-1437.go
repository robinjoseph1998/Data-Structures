package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	start := false
	gap := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if start && gap < k {
				return false
			}
			start = true
			gap = 0
		} else if start && nums[i] == 0 {
			gap++
		}
	}
	return true
}

func main() {
	nums := []int{1, 0, 0, 1, 0, 0, 1}
	k := 2
	fmt.Println(kLengthApart(nums, k))
}
