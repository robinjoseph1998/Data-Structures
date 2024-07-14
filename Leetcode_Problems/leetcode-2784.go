package main

import (
	"fmt"
	"sort"
)

func isGood(nums []int) bool {
	sort.Ints(nums)
	myMap := make(map[int]int)
	for _, v := range nums {
		myMap[v]++
	}
	occ := 0
	expectedlength := nums[len(nums)-1] + 1
	val := myMap[expectedlength-1]
	for _, v := range myMap {
		if v > 1 {
			occ++
		}
	}
	return len(nums) == expectedlength && val == 2 && occ == 1
}

func main() {
	arr := []int{1, 3, 3, 2}
	fmt.Println(isGood(arr))
}
