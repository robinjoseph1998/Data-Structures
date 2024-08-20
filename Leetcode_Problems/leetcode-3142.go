package main

import "fmt"

func satisfiesConditions(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+1 < len(grid) && grid[i][j] != grid[i+1][j] {
				return false
			}

			if j+1 < len(grid[i]) && grid[i][j] == grid[i][j+1] {
				return false
			}
		}
	}
	return true
}

func main() {
	grid := [][]int{{1, 0, 2}, {1, 0, 2}}
	fmt.Println(satisfiesConditions(grid))
}
