package main

import (
	"fmt"
)

func getSneakyNumbers(nums []int) []int {
	myMap := make(map[int]int)
	for _, v := range nums {
		myMap[v]++
	}

	res := []int{}
	for k, v := range myMap {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	nums := []int{0, 1, 1, 0}
	fmt.Println(getSneakyNumbers(nums))
}
