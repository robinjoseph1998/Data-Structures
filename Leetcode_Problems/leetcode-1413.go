package main

import (
	"fmt"
	"math"
)

func minStartValue(nums []int) int {
	res := 0
	for i := 1; i < math.MaxInt; i++ {
		startValue := i
		for j := 0; j < len(nums); j++ {
			startValue += nums[j]
			if startValue < 1 {
				break
			}
			if j == len(nums)-1 {
				res = i
				break
			}
		}
		if res != 0 {
			break
		}
	}
	return res
}

func main() {
	nums := []int{-12}
	fmt.Println(minStartValue(nums))
}
