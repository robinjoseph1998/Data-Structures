package main

import "fmt"

func countHillValley(nums []int) int {
	Helper := make([]int, 0, len(nums))
	Helper = append(Helper, nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			Helper = append(Helper, nums[i])
		}
	}
	count := 0

	for k := 1; k < len(Helper)-1; k++ {
		if Helper[k] > Helper[k-1] && Helper[k] > Helper[k+1] || Helper[k] < Helper[k-1] && Helper[k] < Helper[k+1] {
			count++
		}
	}

	return count
}

func main() {

	nums := []int{21, 21, 21, 2, 2, 2, 2, 21, 21, 45}

	fmt.Println(countHillValley(nums))
}
