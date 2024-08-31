package main

import "fmt"

func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if k == Abs(nums[i], nums[j]) {
				count++
			}
		}
	}
	return count
}

func Abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {

	nums := []int{1, 2, 2, 1}
	k := 1
	fmt.Println(countKDifference(nums, k))

}
