package main

import "fmt"

func isMonotonic(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	increasing := true
	decreasing := true
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			increasing = false
		}
		if nums[i] < nums[i-1] {
			decreasing = false
		}
	}
	return increasing || decreasing
}

func main() {
	nums := []int{1, 3, 2}
	fmt.Println(isMonotonic(nums))
}
