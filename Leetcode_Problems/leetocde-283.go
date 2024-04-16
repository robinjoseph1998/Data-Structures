package main

import "fmt"

func moveZeroes(nums []int) {
	nonZeroIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIdx] = nums[i]
			nonZeroIdx++
		}
	}
	for i := nonZeroIdx; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println("nums", nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
}
