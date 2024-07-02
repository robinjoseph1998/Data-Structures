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
	nums := []int{1, 3, 5}
	fmt.Println(findMin(nums))
}
