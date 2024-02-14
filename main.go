package main

import "fmt"

func distributeCandies(n int, limit int) int {
	count := 0
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			for k := 0; k <= limit; k++ {
				if i+j+k == n {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	n := 5
	limit := 2
	fmt.Println(distributeCandies(n, limit))
}
