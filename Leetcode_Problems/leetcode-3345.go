package main

import (
	"fmt"
	"strconv"
)

func smallestNumber(n int, t int) int {
	for !isDivided(n, t) {
		n++
	}
	return n
}

func isDivided(num, t int) bool {
	strN := strconv.Itoa(num)
	product := 1
	for _, v := range strN {
		digit, _ := strconv.Atoi(string(v))
		product *= digit
	}
	return product%t == 0
}
func main() {
	n := 10
	t := 2
	fmt.Println(smallestNumber(n, t))
}
