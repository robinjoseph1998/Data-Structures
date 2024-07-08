package main

import "fmt"

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix)/2; j++ {
			k := len(matrix) - j - 1
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
	fmt.Println("matrix", matrix)
}

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
}
