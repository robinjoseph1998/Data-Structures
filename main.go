package main

import (
	"fmt"
)

func findMissingAndRepeatedValues(grid [][]int) []int {
	Freq := make(map[int]int)
	for _, each := range grid {
		for _, val := range each {
			Freq[val]++
		}
	}
	n := len(grid) * len(grid)
	var repeat, missing int

	for i := 1; i <= n; i++ {
		if Freq[i] == 2 {
			repeat = i
		}
		if Freq[i] == 0 {
			missing = i
		}
	}
	return []int{repeat, missing}
}

func main() {

	grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(findMissingAndRepeatedValues(grid))

}
