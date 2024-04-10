package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	resMap := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				resMap[nums1[i]]++
			}
		}
	}

	var res []int
	for key := range resMap {
		res = append(res, key)
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	fmt.Println(intersection(nums1, nums2))
}
