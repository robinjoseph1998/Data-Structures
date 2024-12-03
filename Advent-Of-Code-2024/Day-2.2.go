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

func checkingByRmAValue(arr []int) int {
	count := 0
	for i := range arr {
		res := []int{}
		for j, val := range arr {
			if i != j {
				res = append(res, val)
			}

		}
		if isSafe(res) {
			count++
			break
		}
	}
	return count
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {

	str := `27 29 32 33 36 37 40 37
61 62 63 66 66
53 54 57 60 61 62 63 67
5 7 10 12 14 16 18 23
71 74 77 80 83 80 83
22 25 28 30 33 30 32 29
44 47 49 48 48
3 6 8 10 13 14 11 15`

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
	count1, count2 := 0, 0
	for _, v := range nums {
		if isSafe(v) {
			count1++
		} else {
			count2 += checkingByRmAValue(v)
		}
	}
	fmt.Println("count1", count1)
	fmt.Println("count2", count2)

	fmt.Println("total safe:", count1+count2)
}
