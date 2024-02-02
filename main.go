package main

import (
	"fmt"
	"sort"
)

func maximumBeauty(nums []int, k int) int {
	sort.Ints(nums)
	L := 0
	max := 1

	for i := 0; i < len(nums); i++ {
		for nums[L]+k < nums[i]-k {
			L++
		}
		if i-L+1 > max {
			max = i - L + 1
		}
	}
	return max
}

func main() {

	nums := []int{4, 6, 1, 2}
	k := 2
	fmt.Println(maximumBeauty(nums, k))
}
