package main

import (
	"fmt"
	"strconv"
)

func lastVisitedIntegers(words []string) []int {
	var result []int
	var nums []int
	count := 0
	for i := 0; i < len(words); i++ {
		if words[i] != "prev" {
			val, _ := strconv.Atoi(words[i])
			nums = append(nums, val)
			count = 0
		} else {
			count++
			if count <= len(nums) {
				index := len(nums) - count
				result = append(result, nums[index])
			} else {
				result = append(result, -1)
			}
		}
	}
	return result
}

func main() {
	words := []string{"1", "prev", "2", "prev", "prev"}
	fmt.Println(lastVisitedIntegers(words))
}
