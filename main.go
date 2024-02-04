package main

import (
	"fmt"
	"math/bits"
)

func sumIndicesWithKSetBits(nums []int, k int) int {
	sum := 0
	for i, each := range nums {
		Ones := bits.OnesCount(uint(i))
		if Ones == k {
			sum += each
		}
	}
	return sum
}

func main() {
	nums := []int{4, 3, 2, 1}
	k := 1
	fmt.Println(sumIndicesWithKSetBits(nums, k))
}
