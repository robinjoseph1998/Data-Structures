package main

import "fmt"

func occurrencesOfElement(nums []int, queries []int, x int) []int {
	positions := []int{}
	for idx, v := range nums {
		if v == x {
			positions = append(positions, idx)
		}
	}
	res := make([]int, len(queries))
	for i, index := range queries {
		if index-1 < len(positions) {
			res[i] = positions[index-1]
		} else {
			res[i] = -1
		}
	}
	return res
}
func main() {
	nums := []int{1, 3, 1, 7}
	queries := []int{1, 3, 2, 4}
	x := 1
	fmt.Println(occurrencesOfElement(nums, queries, x))
}
