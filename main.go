package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	index := 1
	Ranks := make(map[int]int)

	for _, val := range sortedArr {
		if Ranks[val] == 0 {
			Ranks[val] = index
			index++
		}
	}

	for i, v := range arr {
		arr[i] = Ranks[v]
	}
	return arr
}

func main() {
	arr := []int{37, 12, 28, 9, 100, 56, 80, 5, 12}
	fmt.Println(arrayRankTransform(arr))
}
