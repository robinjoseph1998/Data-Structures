package main

import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	if isIncreasing(heights) {
		return 0
	}
	temp := make([]int, len(heights))
	copy(temp, heights)
	sort.Ints(heights)
	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != temp[i] {
			count++
		}
	}
	return count
}

func isIncreasing(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	heights := []int{2, 1, 2, 1, 1, 2, 2, 1}
	fmt.Println(heightChecker(heights))
}
