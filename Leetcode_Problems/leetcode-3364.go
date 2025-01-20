package main

import (
	"fmt"
	"math"
)

func minimumSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	minSum := math.MaxInt32
	found := false

	for size := l; size <= r; size++ {
		windowSum := 0

		for j := 0; j < size; j++ {
			windowSum += nums[j]
		}

		if windowSum > 0 {
			minSum = min(minSum, windowSum)
			found = true
		}

		for k := size; k < n; k++ {
			windowSum += nums[k] - nums[k-size]
			if windowSum > 0 {
				minSum = min(minSum, windowSum)
				found = true
			}
		}
	}
	if !found {
		return -1
	}
	return minSum
}

func Min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func main() {
	nums := []int{3, -2, 1, 4}
	l := 2
	r := 3
	fmt.Println(minimumSumSubarray(nums, l, r))
}
