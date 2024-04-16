package main

import (
	"fmt"
	"sync"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	var wg sync.WaitGroup
	var res1 []int
	var res2 []int
	wg.Add(2)
	go func(nums1 []int, nums2 []int) {
		defer wg.Done()
		for i := 0; i < len(nums1); i++ {
			found := false
			for j := 0; j < len(nums2); j++ {
				if nums1[i] == nums2[j] {
					found = true
					break
				}
			}
			if !found {
				res1 = appendIfUnique(res1, nums1[i])
			}
		}
	}(nums1, nums2)
	go func(nums1 []int, nums2 []int) {
		defer wg.Done()
		for i := 0; i < len(nums2); i++ {
			found := false
			for j := 0; j < len(nums1); j++ {
				if nums2[i] == nums1[j] {
					found = true
					break
				}
			}
			if !found {
				res2 = appendIfUnique(res2, nums2[i])
			}
		}
	}(nums1, nums2)
	wg.Wait()
	return [][]int{res1, res2}
}

func appendIfUnique(slice []int, element int) []int {
	for _, val := range slice {
		if val == element {
			return slice
		}
	}
	return append(slice, element)
}
func main() {
	var wg sync.WaitGroup
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	fmt.Println(findDifference(nums1, nums2))
	wg.Wait()
}
