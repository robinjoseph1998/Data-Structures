package main

import "fmt"

func mostFrequent(nums []int, key int) int {
	var (
		myMap  = make(map[int]int)
		max    = 0
		result = 0
	)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			myMap[nums[i+1]]++
			if myMap[nums[i+1]] > max {
				max = myMap[nums[i+1]]
				result = nums[i+1]
			}
		}
	}
	return result
}

func main() {
	nums := []int{1, 100, 200, 1, 100}
	key := 1
	fmt.Println(mostFrequent(nums, key))
}
