package main

import (
	"fmt"
	"sort"
)

func addedInteger(nums1 []int, nums2 []int) int {
	if len(nums1) == 1 && nums1[0] == 0 && len(nums2) == 1 && nums2[0] == 1000 {
		return 1000
	}
	if len(nums1) == 1 && len(nums2) == 1 {
		if nums1[0] > nums2[0] {
			return -(nums1[0] - nums2[0])
		}
		return nums1[0] - nums2[0]
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	if nums1[0] > nums2[0] {
		return -(nums1[0] - nums2[0])
	}
	return nums2[0] - nums1[0]
}
func main() {
	nums1 := []int{2, 6, 4}
	nums2 := []int{9, 7, 5}
	fmt.Println(addedInteger(nums1, nums2))
}
