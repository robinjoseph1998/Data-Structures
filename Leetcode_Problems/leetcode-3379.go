package main

import "fmt"

func constructTransformedArray(nums []int) []int {
	result := []int{}
	i := 0
	for i < len(nums) {
		pos := (i + nums[i]%len(nums) + len(nums)) % len(nums)
		result = append(result, nums[pos])
		i++
	}
	return result
}

func main() {
	nums := []int{3, -2, 1, 1}
	fmt.Println(constructTransformedArray(nums))
}
