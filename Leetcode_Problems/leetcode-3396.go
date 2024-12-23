package main

import "fmt"

func minimumOperations(nums []int) int {
	operations := 0
	for len(nums) > 0 {
		if !isDistinct(nums) {
			operations++
			if len(nums) > 3 {
				nums = nums[3:]
			} else {
				nums = []int{}
			}
		} else {
			break
		}
	}
	return operations
}

func isDistinct(nums []int) bool {
	seen := make(map[int]bool)
	for _, val := range nums {
		if seen[val] {
			return false
		}
		seen[val] = true
	}
	return true
}
func main() {
	nums := []int{4, 5, 6, 4, 4}
	fmt.Println(minimumOperations(nums))
}
