package main

import (
	"fmt"
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	appleTotal := 0
	for _, v := range apple {
		appleTotal += v
	}
	bags := 0
	count := 0
	res := 0
	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	for _, v := range capacity {
		bags += v
		count++
		if bags >= appleTotal {
			res = count
			break
		}
	}
	return res
}

func main() {
	apple := []int{1, 8, 3, 3, 5}
	capacity := []int{3, 9, 5, 1, 9}
	fmt.Println(minimumBoxes(apple, capacity))
}
