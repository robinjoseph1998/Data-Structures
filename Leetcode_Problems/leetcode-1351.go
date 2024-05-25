package main

import "fmt"

func countNegatives(grid [][]int) int {
	count := 0
	for _, each := range grid {
		for _, num := range each {
			if num < 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	fmt.Println(countNegatives(grid))
}
