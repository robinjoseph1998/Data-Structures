package main

import "fmt"

func findChampion(grid [][]int) int {
	count, big, idx := 0, 0, 0
	for dx, each := range grid {
		for i := 0; i < len(each); i++ {
			if each[i] == 1 {
				count++
			}
		}
		if count > big {
			big = count
			idx = dx
			fmt.Println("dx", dx)
		}
		count = 0
	}
	return idx
}

func main() {

	grid := [][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}}
	fmt.Println(findChampion(grid))

}
