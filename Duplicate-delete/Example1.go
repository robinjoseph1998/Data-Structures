package main

import (
	"fmt"
)

func removeDuplicates(slice []int) []int {
	occ := make(map[int]int)
	for _, v := range slice {
		occ[v]++
	}

	slice = []int{}
	for v := range occ {
		slice = append(slice, v)
	}
	return slice
}

func main() {
	slice := []int{1, 2, 3, 2, 4, 5, 1, 6}
	fmt.Println("Original Slice:", slice)
	fmt.Println("Slice without duplicates:", removeDuplicates(slice))
}
