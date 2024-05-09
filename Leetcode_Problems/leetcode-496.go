package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := []int{}
	aff := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				for k := j + 1; k < len(nums2); k++ {
					if nums2[k] > nums1[i] {
						res = append(res, nums2[k])
						aff = 1
						break
					}
				}
				if aff == 0 {
					res = append(res, -1)
				}
				aff = 0
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))

}
