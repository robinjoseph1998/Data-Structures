package main

import "fmt"

func divide(d int, di int) int {
	if d == -2147483648 && di == -1 {
		return 2147483647
	}
	return d / di
}

func main() {
	dividend := -2147483648
	divisor := -1
	fmt.Println(divide(dividend, divisor))
}
