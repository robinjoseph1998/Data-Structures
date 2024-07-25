package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)
	res := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if (len(cost)-i)%3 != 0 {
			res += cost[i]
		}
	}
	return res
}

func main() {
	cost := []int{6, 5, 7, 9, 2, 2}
	fmt.Println(minimumCost(cost))
}
