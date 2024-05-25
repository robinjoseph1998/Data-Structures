package main

import (
	"fmt"
	"math"
)

func shortestToChar(s string, c byte) []int {
	positions := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			positions = append(positions, i)
		}
	}
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		minDistance := math.MaxInt64
		for _, pos := range positions {
			distance := abs(i - pos)
			if distance < minDistance {
				minDistance = distance
			}
		}
		res[i] = minDistance
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	s := "loveleetcode"
	c := "e"
	fmt.Println(shortestToChar(s, c[0]))
}
