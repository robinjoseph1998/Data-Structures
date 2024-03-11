package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	var arr1 []int
	for i := 0; i < m; i++ {
		arr1 = append(arr1, nums1[i])
	}
	for i := 0; i < n; i++ {
		arr1 = append(arr1, nums2[i])
	}
	copy(nums1, arr1)
	sort.Ints(nums1)
	fmt.Println("nums1[i]", nums1)
}

func main() {

	nums1 := []int{-50, -48, -47, -47, -46, -44, -43, -43, -41, -39, -38, -37, -37, -37, -33, -33, -32, -31, -29, -29, -27, -26, -26, -26, -25, -25, -24, -24, -23, -22, -22, -22, -18, -17, -17, -14, -14, -11, -11, -11, -6, -5, -5, -5, -5, -4, -1, 0, 2, 2, 2, 2, 5, 6, 7, 7, 9, 10, 13, 13, 14, 14, 18, 21, 21, 21, 22, 24, 24, 24, 25, 26, 26, 29, 29, 29, 30, 30, 31, 31, 32, 33, 34, 34, 34, 35, 38, 40, 41, 42, 43, 44, 44, 46, 46, 47, 47, 48, 49, 49}
	m := 100
	nums2 := []int{}

	n := 0
	merge(nums1, m, nums2, n)

}
