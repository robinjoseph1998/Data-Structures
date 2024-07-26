package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	fmt.Println("nums", nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		fmt.Println("nums[i]", nums[i])
		res += nums[i]
	}
	return res
}

func main() {
	nums := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(nums))
}
