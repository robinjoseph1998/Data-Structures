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
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	fmt.Println(searchMatrix(matrix, target))
}
