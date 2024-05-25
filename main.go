package main

import (
	"fmt"
	"math"
)

func convertToBase7(num int) string {
	res := math.MaxInt64

	for res > 0 {
		res = num / 7
	}
	fmt.Println("res", res)
	return ""
}

func main() {
	num := 100
	fmt.Println(convertToBase7(num))
}
