package main

import (
	"fmt"
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	i := 0
	var Out [][]int
	for i < len(nums) {
		if i+2 < len(nums) && nums[i+2]-nums[i] <= k {
			Out = append(Out, []int{nums[i], nums[i+1], nums[i+2]})
		} else {
			return [][]int{}
		}
		i += 3
	}
	return Out
}

func main() {
	nums := []int{1, 3, 4, 8, 7, 9, 3, 5, 1}
	k := 2
	fmt.Println(divideArray(nums, k))

}
