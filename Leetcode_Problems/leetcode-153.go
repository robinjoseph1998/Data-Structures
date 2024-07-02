package main

import (
	"fmt"
	"sort"
)

func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
