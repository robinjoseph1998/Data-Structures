package main

import "fmt"

func repeatedNTimes(nums []int) int {
	occ := make(map[int]int)
	for _, v := range nums {
		occ[v]++

	}
	res := 0
	ans := 0
	for k, v := range occ {
		if v > res {
			res = v
			ans = k
		}
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3, 3}
	fmt.Println(repeatedNTimes(nums))
}
