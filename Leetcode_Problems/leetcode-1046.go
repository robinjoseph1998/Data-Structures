package main

import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		last := stones[len(stones)-1]
		secondLast := stones[len(stones)-2]

		stones = stones[:len(stones)-2]
		if last != secondLast {
			stones = append(stones, last-secondLast)
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones))
}
