package main

import "fmt"

func decompressRLElist(nums []int) []int {
	res := []int{}
	for j := 1; j < len(nums); j += 2 {
		for i := 0; i < nums[j-1]; i++ {
			res = append(res, nums[j])
		}
	}
	return res
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums))
}
