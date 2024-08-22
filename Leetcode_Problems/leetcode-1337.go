package main

import (
	"fmt"
	"sort"
)

type Pairs struct {
	index   int
	soilder int
}

func kWeakestRows(mat [][]int, k int) []int {
	rows := []Pairs{}
	for idx, row := range mat {
		soilders := 0
		for _, v := range row {
			if v == 0 {
				break
			}
			soilders += 1
		}
		rows = append(rows, Pairs{index: idx, soilder: soilders})
	}

	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].soilder < rows[j].soilder
	})
	res := []int{}

	for i := 0; i < k; i++ {
		res = append(res, rows[i].index)
	}

	return res
}
func main() {

	mat := [][]int{{1, 0}, {0, 0}, {1, 0}, {1, 1}, {1, 1}, {1, 1}, {1, 1}}
	k := 4

	fmt.Println(kWeakestRows(mat, k))

}
