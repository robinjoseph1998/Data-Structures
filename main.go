package main

import (
	"fmt"
	"math"
)

func findClosestNumber(nums []int) int {
	close := math.MaxInt64

	for _, num := range nums {
		if math.Abs(float64(num)) < math.Abs(float64(close)) ||
			math.Abs(float64(num)) == math.Abs(float64(close)) && num > close {
			close = num

		}
	}
	return close
}

func main() {
	nums := []int{-4, -2, 1, 4, 8}
	fmt.Println(findClosestNumber(nums))
}
