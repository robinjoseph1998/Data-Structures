package main

import "fmt"

func search(nums []int, target int) bool {
	for _, val := range nums {
		if target == val {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
