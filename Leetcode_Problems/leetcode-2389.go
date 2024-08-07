package main

import (
	"fmt"
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	res := make([]int, len(queries))
	for i, queryNum := range queries {
		tot := 0
		count := 0
		for j := 0; j < len(nums); j++ {
			if tot+nums[j] <= queryNum {
				tot += nums[j]
				count++
			} else {
				break
			}
		}
		res[i] = count
	}
	return res
}

func main() {
	nums := []int{4, 5, 2, 1}
	queries := []int{3, 10, 21}
	fmt.Println(answerQueries(nums, queries))
}
