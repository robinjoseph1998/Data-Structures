package main

import "fmt"

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	c := 1
	for n >= c {
		n -= c
		c++
	}
	return c - 1
}

func main() {
	n := 5
	fmt.Println(arrangeCoins(n))
}
