package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	tot := 0
	for i := 5; i <= n; i *= 5 {
		tot += n / i
	}
	return tot
}

func main() {
	n := 30
	fmt.Println(trailingZeroes(n))
}
