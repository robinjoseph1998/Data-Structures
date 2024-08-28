package main

import (
	"fmt"
	"sort"
)

func findGCD(nums []int) int {
	sort.Ints(nums)

	return gcd(nums[0], nums[len(nums)-1])
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func main() {
	nums := []int{2, 5, 6, 9, 10}
	fmt.Println(findGCD(nums))
}
