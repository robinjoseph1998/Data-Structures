package main

import "fmt"

func search(nums []int, target int) int {
	for idx, val := range nums {
		if val == target {
			return idx
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
