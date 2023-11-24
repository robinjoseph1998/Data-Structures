package main

import "fmt"

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	var BigNums [][]int
	BigNums = append(BigNums, nums1, nums2, nums3)
	i := 0
	var res []int
	var result []int
	for i <= len(BigNums)-1 {
		arr := BigNums[i]
		helpMap := make(map[int]bool)
		for _, val := range arr {
			if !helpMap[val] {
				helpMap[val] = true
				result = append(result, val)
			}
		}
		i++
	}
	ResultMap := make(map[int]int)
	for _, val := range result {
		ResultMap[val]++
	}
	for key, val := range ResultMap {
		if val > 1 {
			res = append(res, key)
		}
	}
	return res
}
func main() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}

	fmt.Println(twoOutOfThree(nums1, nums2, nums3))
}
