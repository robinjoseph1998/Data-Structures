package main

import "fmt"

func countOdds(low int, high int) int {
	count := 0
	for low <= high {
		if low%2 != 0 {
			count++
		}
		low++
	}
	return count
}

func main() {
	low := 3
	high := 7
	fmt.Println(countOdds(low, high))
}
