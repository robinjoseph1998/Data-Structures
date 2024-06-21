package main

import "fmt"

func modifiedMatrix(matrix [][]int) [][]int {
	negMap := make(map[int][]int)
	for arIdx, M := range matrix {
		for idx, E := range M {
			if E < 0 {
				negMap[arIdx] = append(negMap[arIdx], idx)
			}
		}
	}
	tmp := 0
	for k, v := range negMap {
		for _, val := range v {
			for _, m := range matrix {
				if m[val] > tmp {
					tmp = m[val]
				}
			}
			matrix[k][val] = tmp
			tmp = 0
		}
	}

	return matrix
}

func main() {
	matrix := [][]int{{-1, 0, 0, 2, 2}, {2, 0, 0, 2, 1}, {4, 3, 2, 1, 1}, {-1, -1, 0, 2, 4}, {1, 0, 3, -1, 0}}
	fmt.Println(modifiedMatrix(matrix))
}
