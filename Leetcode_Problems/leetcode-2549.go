package main

import "fmt"

func distinctIntegers(n int) int {
	board := make(map[int]bool)
	board[n] = true
	Added := true
	for Added {
		Added = false
		currentNums := []int{}
		for v := range board {
			currentNums = append(currentNums, v)
		}
		for _, v := range currentNums {
			for i := 1; i <= n; i++ {
				if v%i == 1 && !board[i] {
					board[i] = true
					Added = true
				}
			}
		}
	}
	return len(board)
}

func main() {
	n := 5
	fmt.Println(distinctIntegers(n))
}
