package main

import (
	"fmt"
	"math"
	"sort"
)

func findMaxK(nums []int) int {
	sort.Ints(nums)
	fmt.Println("nums", nums)
	var result int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < 0 {
				pInt := int(math.Abs(float64(nums[i])))
				if pInt == nums[j] {
					fmt.Println("pInt", pInt)
					if result < pInt {
						result = pInt
					}
				} else {
					continue
				}
			}
		}
	}
	if result != 0 {
		return result
	}
	return -1
}

func main() {

	nums := []int{-1, 10, 6, 7, -7, 1}
	fmt.Println(findMaxK(nums))
}
