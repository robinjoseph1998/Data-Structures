package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var rightDiagonal int32
	var leftDiagonal int32

	for i := 0; i < len(arr); i++ {
		rightDiagonal += arr[i][i]
		leftDiagonal += arr[i][len(arr)-i-1]
	}
	return abs(rightDiagonal, leftDiagonal)
}

func abs(a, b int32) int32 {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	arr := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	fmt.Println(diagonalDifference(arr))
}
