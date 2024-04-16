package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	first := make(map[int]int)
	second := make(map[int]int)
	for _, val := range nums1 {
		first[val]++
	}

	for _, val := range nums2 {
		second[val]++
	}
	for key, val := range first {
		if v, ok := second[key]; ok {
			count := minCount(val, v)
			for i := 0; i < count; i++ {
				res = append(res, key)
			}
		}
	}
	return res
}

func minCount(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	nums1 := []int{54, 93, 21, 73, 84, 60, 18, 62, 59, 89, 83, 89, 25, 39, 41, 55, 78, 27, 65, 82, 94, 61, 12, 38, 76, 5, 35, 6, 51, 48, 61, 0, 47, 60, 84, 9, 13, 28, 38, 21, 55, 37, 4, 67, 64, 86, 45, 33, 41}
	nums2 := []int{17, 17, 87, 98, 18, 53, 2, 69, 74, 73, 20, 85, 59, 89, 84, 91, 84, 34, 44, 48, 20, 42, 68, 84, 8, 54, 66, 62, 69, 52, 67, 27, 87, 49, 92, 14, 92, 53, 22, 90, 60, 14, 8, 71, 0, 61, 94, 1, 22, 84, 10, 55, 55, 60, 98, 76, 27, 35, 84, 28, 4, 2, 9, 44, 86, 12, 17, 89, 35, 68, 17, 41, 21, 65, 59, 86, 42, 53, 0, 33, 80, 20}

	fmt.Println(intersect(nums1, nums2))
}
