package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	for _, v := range matrix {
		for _, num := range v {
			if num == target {
				return true
			}
		}
	}
	return false
}
func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 3
	fmt.Println(searchMatrix(matrix, target))
}
