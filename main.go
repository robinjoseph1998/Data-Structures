package main

import (
	"fmt"
	"sort"
)

func sumOfPower(nums []int) int {
	modulo := 1_000_000_007
	sort.Ints(nums)
	sum, result := 0, 0
	for _, num := range nums {
		result = (result + (num*num%modulo)*(num+sum)) % modulo
		sum = (sum<<1 + num) % modulo

	}
	return result
}

func main() {
	nums := []int{2, 1, 4}
	fmt.Println(sumOfPower(nums))
}
