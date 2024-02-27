package main

import (
	"fmt"
	"sort"
)

func numberGame(nums []int) []int {
	var arr []int
	sort.Sort(sort.Reverse(sort.Reverse(sort.IntSlice(nums))))
	fmt.Println("nums", nums)
	for i := 0; i < len(nums); i += 2 {
		if i+1 < len(nums) && nums[i] < nums[i+1] {
			if nums[i] < nums[i+1] {
				arr = append(arr, nums[i+1])
				arr = append(arr, nums[i])
			}
		}
		if nums[i] == nums[i+1] {
			arr = append(arr, nums[i])
			arr = append(arr, nums[i+1])
		}
	}
	return arr
}

func main() {

	nums := []int{2, 7, 9, 6, 4, 6}
	fmt.Println(numberGame(nums))

}
