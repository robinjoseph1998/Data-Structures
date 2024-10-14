package main

import "fmt"

func findSubarrays(nums []int) bool {
	sumMap := make(map[int]int)
	for i := 1; i < len(nums); i++ {
		sum := nums[i] + nums[i-1]
		sumMap[sum]++

		if sumMap[sum] == 2 {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(findSubarrays(nums))
}
