package main

import (
	"fmt"
	"math"
)

func isStrictlyIncreasing(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func primeSubOperation(nums []int) bool {
	if isStrictlyIncreasing(nums) {
		return true
	}
	var Helper []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 2 {
			for j := nums[i] - 1; j > 0; j-- {
				for p := 2; p < int(math.Sqrt(float64(j))); p++ {
					if j%p != 0 {
						Helper = append(Helper, j)
					}
				}
			}
		}
	}
	fmt.Println("Helper", Helper)
	return false
}

func main() {
	nums := []int{4, 9, 6, 10}

	fmt.Println(primeSubOperation(nums))
}
