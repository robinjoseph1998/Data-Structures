package main

import "fmt"

func xorOperation(n int, start int) int {
	res := 0
	for i := 0; i < n; i++ {
		res ^= start + 2*i
	}

	return res
}

func main() {
	n := 5
	start := 0
	fmt.Println(xorOperation(n, start))
}
