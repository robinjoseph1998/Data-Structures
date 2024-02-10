package main

import "fmt"

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			val1, val2 := abs(i, j, nums, indexDifference, valueDifference)
			if val1 != -9 && val2 != -9 {
				result = append(result, val1)
				result = append(result, val2)
			}
		}
	}
	fmt.Println("result", result)
	if len(result) == 0 {
		result = append(result, -1)
		result = append(result, -1)
	}
	return result
}

func abs(i int, j int, nums []int, indexDifference int, valueDifference int) (a, b int) {
	max := Min(nums[i], nums[j])
	min := Max(nums[i], nums[j])
	if j-i >= indexDifference && min-max >= valueDifference {
		return i, j
	}
	return -9, -9
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{31, 23, 36}
	indexDifference := 1
	valueDifference := 11
	fmt.Println(findIndices(nums, indexDifference, valueDifference))
}
