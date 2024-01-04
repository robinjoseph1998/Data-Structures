package main

import (
	"fmt"
)

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	var sum int
	for i, val := range nums1 {
		diff := val - nums2[i]
		sum += diff
	}
	if sum%k != 0 {
		return -1
	}
	var Moves int64
	for i := 0; i < len(nums1); i++ {
		diff := nums1[i] - nums2[i]

		if diff < 0 {
			Moves += int64((diff + k - 1) / k)
			nums1[i] += k
		}

		if diff > 0 {
			Moves += int64(diff / k)
			nums1[i] -= k
		}

	}
	return Moves
}

func main() {
	nums1 := []int{4, 3, 1, 4}
	nums2 := []int{1, 3, 7, 1}
	k := 1
	fmt.Println(minOperations(nums1, nums2, k))

}
