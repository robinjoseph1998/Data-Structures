package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isSafe(nums []int) bool {
	if nums[0] > nums[1] {
		if isDecreasing(nums) {
			return true
		}
	}
	if nums[0] < nums[1] {
		if isIncreasing(nums) {
			return true
		}
	}
	return false
}

func isIncreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] || abs(nums[i], nums[i-1]) > 3 || nums[i] == nums[i-1] {
			return false
		}
	}
	return true
}

func isDecreasing(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] || abs(nums[i], nums[i-1]) > 3 || nums[i] == nums[i-1] {
			return false
		}
	}
	return true
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {

	str := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	rows := strings.Split(str, "\n")

	var nums [][]int

	for _, each := range rows {
		numStr := strings.Fields(each)
		arr := []int{}
		for _, num := range numStr {
			dig, _ := strconv.Atoi(num)
			arr = append(arr, dig)
		}
		nums = append(nums, arr)
	}
	count := 0
	for _, v := range nums {
		if isSafe(v) {
			count++
		}
	}
	fmt.Println("count", count)
}
