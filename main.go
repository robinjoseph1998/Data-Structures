package main

import (
	"fmt"
	"sort"
)

func minimumCost(nums []int) int {
	ans := nums[0]
	sort.Ints(nums[1:])
	fmt.Println("nums", nums)
	return ans + nums[1] + nums[2]
}

func main() {
	nums := []int{1, 2, 3, 12}
	fmt.Println(minimumCost(nums)) // Output: 6

}
