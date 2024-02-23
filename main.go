package main

import (
	"fmt"
	"sort"
)

func findMissingAndRepeatedValues(grid [][]int) []int {
	var arr []int
	for _, each := range grid {
		arr = append(arr, each...)
	}
	sort.Ints(arr)
	Map := make(map[int]int)
	for _, val := range arr {
		Map[val]++
	}
	var res []int
	for val, freq := range Map {
		if freq > 1 {
			res = append(res, val)
		}
	}
	missing := 0
	count := 0
	fmt.Println("arr", arr)
	for i := arr[0]; i <= len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == arr[j] {
				count = 0
			} else {
				count++
			}
		}
		if i == len(arr) && count == 1 {
			missing = i
		}
	}
	res = append(res, missing)
	return res
}

func main() {

	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(findMissingAndRepeatedValues(grid))

}
