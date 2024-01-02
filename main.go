package main

import (
	"fmt"
	"strconv"
)

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	var str string
	var diff []int
	for i := 0; i < len(nums2); i++ {
		str = strconv.Itoa(nums2[i])

	}
	fmt.Println("str", str)
	// intstr, _ := strconv.Atoi(str)
	// Count := 0
	for i := 0; i < len(nums1); i++ {
		if i%2 == 0 {
			if nums1[i] != nums2[i] {
				diff = append(diff, nums1[i]+k)
			}
		} else {
			if nums1[i] != nums2[i] {
				diff = append(diff, nums1[i]-k)
			}
		}
	}
	fmt.Println("diff", diff)
	return 0
}

func main() {
	nums1 := []int{4, 3, 1, 4}
	nums2 := []int{1, 3, 7, 1}
	k := 3
	fmt.Println(minOperations(nums1, nums2, k))

}
