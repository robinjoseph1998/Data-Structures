package main

import (
	"fmt"
)

func prefixesDivBy5(nums []int) []bool {
	num := 0
	res := make([]bool, len(nums))
	for i, v := range nums {
		num = (num<<1 | v) % 5
		res[i] = (num == 0)
	}
	return res
}
func main() {
	dominoes := []int{0, 1, 1, 1, 1, 1}
	fmt.Println(prefixesDivBy5(dominoes))
}
