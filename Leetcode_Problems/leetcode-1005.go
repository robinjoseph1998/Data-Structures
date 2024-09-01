package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 != 0 {
		sort.Ints(nums)
		nums[0] = -nums[0]
	}
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	nums := []int{4, 2, 3}
	k := 1
	fmt.Println(largestSumAfterKNegations(nums, k))
}
