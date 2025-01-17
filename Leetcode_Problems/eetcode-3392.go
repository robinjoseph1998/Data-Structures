package main

import "fmt"

func countSubarrays(nums []int) int {
	result := 0
	for i := 1; i < len(nums)-1; i++ {
		if (nums[i-1]+nums[i+1])*2 == nums[i] {
			result++
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 1, 4, 1}
	fmt.Println(countSubarrays(nums))
}
