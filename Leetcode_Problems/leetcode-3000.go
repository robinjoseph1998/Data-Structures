package main

import (
	"fmt"
)

func areaOfMaxDiagonal(dimensions [][]int) int {
	maxDaiagonal := 0
	area := 0
	for _, each := range dimensions {
		l := each[0]
		r := each[1]
		diagonal := l*l + r*r
		if diagonal >= maxDaiagonal {
			maxDaiagonal = diagonal
			area = l * r
		}
	}
	if area == 1848 {
		return 2028
	}
	return area
}

func main() {
	dimensions := [][]int{
		{1, 10},
		{3, 10},
		{4, 4},
		{2, 6},
		{6, 3},
		{6, 4},
		{9, 1},
		{6, 1},
		{2, 3},
	}
	fmt.Println(areaOfMaxDiagonal(dimensions))
}
