package main

import "fmt"

func maximumCount(nums []int) int {
	pos, neg := 0, 0
	for _, num := range nums {
		if num < 0 {
			pos++
		} else if num > 0 {
			neg++
		}
	}
	return Max(pos, neg)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2}
	fmt.Println(maximumCount(nums))
}
