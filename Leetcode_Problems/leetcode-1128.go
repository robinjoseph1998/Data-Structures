package main

import (
	"fmt"
)

func numEquivDominoPairs(dominoes [][]int) int {
	count := 0
	PairMap := make(map[[2]int]int)
	for _, sub := range dominoes {
		if sub[0] > sub[1] {
			sub[0], sub[1] = sub[1], sub[0]
		}
		arr := [2]int{sub[0], sub[1]}
		count += PairMap[arr]
		PairMap[arr]++
	}
	return count
}
func main() {
	dominoes := [][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}
	fmt.Println(numEquivDominoPairs(dominoes))
}
