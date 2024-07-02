package main

import "fmt"

func rob(nums []int) int {
	maxNow := 0
	lastMax := 0
	for _, v := range nums {
		newMax := Max(v+lastMax, maxNow)
		lastMax = maxNow
		maxNow = newMax

	}
	return maxNow
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 3, 1}
	fmt.Println(rob(nums))
}
