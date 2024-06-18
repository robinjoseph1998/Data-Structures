package main

import "fmt"

func duplicateNumbersXOR(nums []int) int {
	countMap := make(map[int]int)
	for _, val := range nums {
		countMap[val]++
	}
	res := 0
	for key, val := range countMap {
		if val == 2 {
			res ^= key
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 1, 3}
	fmt.Println(duplicateNumbersXOR(nums))
}
