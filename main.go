package main

import "fmt"

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var res []int
	count := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				count++
				break
			}
		}
	}
	fmt.Println("res", res)
	return []int{}
}

func main() {
	nums1 := []int{4, 3, 2, 3, 1}
	nums2 := []int{2, 2, 5, 2, 3, 6}

	fmt.Println(findIntersectionValues(nums1, nums2))
}
