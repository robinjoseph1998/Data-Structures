package main

import (
	"fmt"
	"sort"
)

func findClosestElements(arr []int, k int, x int) []int {
	type Pairs struct {
		value      int
		difference int
	}
	var PairSlice []Pairs
	for _, val := range arr {
		diff := abs(val - x)
		PairSlice = append(PairSlice, Pairs{value: val, difference: diff})
	}

	sort.Slice(PairSlice, func(i, j int) bool {
		if PairSlice[i].difference == PairSlice[j].difference {
			return PairSlice[i].value < PairSlice[j].value
		}
		return PairSlice[i].difference < PairSlice[j].difference
	})

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, PairSlice[i].value)
	}
	sort.Ints(result)
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func main() {
	arr := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	k := 3
	x := 5
	fmt.Println(findClosestElements(arr, k, x))
}
