package main

import "fmt"

func construct2DArray(original []int, m int, n int) [][]int {
	var TwoD [][]int
	var k int
	p := 1
	for i := 0; i < m; i++ {
		var row []int
		for j := 0 + k; j < n*p; j++ {
			row = append(row, original[j])
		}
		k++
		p = n
		TwoD = append(TwoD, row)
	}
	return TwoD
}
func main() {
	m := 2
	n := 2
	orginal := []int{1, 2, 3, 4}
	fmt.Println(construct2DArray(orginal, m, n))

}
