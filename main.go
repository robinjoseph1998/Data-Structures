package main

import (
	"fmt"
)

func firstBadVersion(n int) int {
	r, l := n, 0
	mid := r + l/2
	for mid > 0 || mid < n {
		if isBadVeion(mid) {
			return mid
		}
	}

}
func main() {
	nums := []int{3, 3, 3}
	target := 3
	fmt.Println(searchRange(nums, target))

}
