package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		start := nums[i]
		for i < len(nums)-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		end := nums[i]
		if start == end {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
		}

	}
	return result
}

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))

}
