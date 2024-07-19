package main

import "fmt"

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0
	for _, v := range position {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return min(odd, even)
}

func main() {
	position := []int{1, 2, 3}
	fmt.Println(minCostToMoveChips(position))
}
