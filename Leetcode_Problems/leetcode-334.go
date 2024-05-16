package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	num1 := math.MaxInt64
	num2 := math.MaxInt64
	for _, v := range nums {
		if num1 >= v {
			num1 = v
		} else if num2 >= v {
			num2 = v
		} else {
			return true
		}
	}
	return false
}

// brute force approach but time limit exceed
// func increasingTriplet(nums []int) bool {
// 	res := false
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i] < nums[j] && nums[j] < nums[k] {
// 					res = true
// 					break
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

func main() {
	nums := []int{20, 100, 10, 12, 5, 13}
	fmt.Println(increasingTriplet(nums))
}
