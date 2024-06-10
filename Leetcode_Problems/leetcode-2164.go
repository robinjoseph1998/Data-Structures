package main

import (
	"fmt"
	"sort"
)

func sortEvenOdd(nums []int) []int {
	oddValues := []int{}
	evenValues := []int{}
	for idx, v := range nums {
		if idx%2 == 0 {
			evenValues = append(evenValues, v)
		} else {
			oddValues = append(oddValues, v)
		}
	}
	sort.Slice(oddValues, func(i, j int) bool {
		return oddValues[i] > oddValues[j]
	})
	sort.Ints(evenValues)
	res := make([]int, len(nums))
	oddIdx, evenIdx := 0, 0
	for idx := range nums {
		if idx%2 == 0 {
			res[idx] = evenValues[evenIdx]
			evenIdx++
		} else {
			res[idx] = oddValues[oddIdx]
			oddIdx++
		}
	}
	return res
}

func main() {
	nums := []int{36, 45, 32, 31, 15, 41, 9, 46, 36, 6, 15, 16, 33, 26, 27, 31, 44, 34}
	fmt.Println(sortEvenOdd(nums))
}
