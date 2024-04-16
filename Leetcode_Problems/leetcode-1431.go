package main

import (
	"fmt"
	"sort"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	OC := make([]int, len(candies))
	copy(OC, candies)
	sort.Ints(candies)
	big := candies[len(candies)-1]
	var Res []bool
	for _, val := range OC {
		if val+extraCandies >= big {
			Res = append(Res, true)
		} else {
			Res = append(Res, false)
		}
	}

	return Res
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
