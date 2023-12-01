package main

// import (
// 	"fmt"
// 	"sort"
// )

// func maxSubsequence(nums []int, k int) []int {
// 	IndiceToRemove := make(map[int]bool)
// 	var OrginalNums []int
// 	OrginalNums = append(OrginalNums, nums...)

// 	sort.Ints(nums)
// 	fmt.Println("nums", nums)
// 	startIndex := len(nums) - k
// 	FirstK := nums[:startIndex]
// 	fmt.Println("FirstK", FirstK)
// 	for _, val := range FirstK {
// 		IndiceToRemove[val] = true
// 	}
// 	var result []int
// 	for _, val := range OrginalNums {

// 		if !IndiceToRemove[val] {
// 			result = append(result, val)
// 		}
// 	}

// 	return result

// }

// func main() {
// 	nums := []int{3, 4, 3, 3}
// 	k := 2

// 	fmt.Println(maxSubsequence(nums, k))

// }
