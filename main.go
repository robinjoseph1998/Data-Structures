package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	if len(nums) == 2 {
		return nums
	}
	occurence := make(map[int]int)
	for _, num := range nums {
		occurence[num]++
	}
	res := make([]int, 2)
	i := 0
	for k, v := range occurence {
		if v == 1 {
			res[i] = k
			i++
		}
	}
	return res
}
func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}
