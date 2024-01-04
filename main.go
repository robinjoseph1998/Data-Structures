package main

import (
	"fmt"
	"reflect"
)

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	if k == 0 {
		if reflect.DeepEqual(nums1, nums2) {
			return 0
		}
		return -1
	}

	var positive, negative int

	for i := 0; i < len(nums1); i++ {
		diff := nums1[i] - nums2[i]

		if diff%k != 0 {
			return -1
		}

		if diff > 0 {
			positive += diff / k
		} else if diff < 0 {
			negative -= diff / k
		} else {
			continue
		}
	}
	if positive != negative {
		return -1
	}
	return int64(positive)
}

func main() {
	nums1 := []int{4, 3, 1, 4}
	nums2 := []int{1, 3, 7, 1}
	k := 1
	fmt.Println(minOperations(nums1, nums2, k))

}
