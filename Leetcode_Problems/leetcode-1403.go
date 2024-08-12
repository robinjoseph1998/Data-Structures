package main

import (
	"fmt"
	"sort"
)

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sum := 0
	res := []int{}
	for _, v := range nums {
		sum += v
	}
	currSum := 0
	for i := len(nums) - 1; i >= 0; i-- {
		currSum += nums[i]
		res = append(res, nums[i])
		if currSum > sum-currSum {
			break
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 10, 9, 8}
	fmt.Println(minSubsequence(nums))
}
