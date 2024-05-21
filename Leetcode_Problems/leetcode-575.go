package main

import "fmt"

func distributeCandies(candyType []int) int {
	UniqueCandies := make(map[int]bool)
	for _, candy := range candyType {
		UniqueCandies[candy] = true
	}
	if len(UniqueCandies) > len(candyType)/2 {
		return len(candyType) / 2
	}
	return len(UniqueCandies)
}

func main() {
	candyType := []int{1, 1, 2, 2, 3, 3}
	fmt.Println(distributeCandies(candyType))
}
