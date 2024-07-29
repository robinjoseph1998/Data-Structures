package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		strv := strconv.Itoa(v)
		if len(strv)%2 == 0 {
			count++
		}
	}
	return count
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
