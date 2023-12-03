package main

import (
	"fmt"
	"sort"
)

func countElements(nums []int) int {
	sort.Ints(nums)
	var count int
	First := nums[0]
	Last := nums[len(nums)-1]
	for _, val := range nums {
		if First != val && Last != val {
			count++
		}
	}
	return count
}

func main() {
	nums := []int{-3, 3, 3, 90}
	fmt.Println(countElements(nums))
}
