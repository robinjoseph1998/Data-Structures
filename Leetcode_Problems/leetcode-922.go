package main

import "fmt"

func sortArrayByParityII(nums []int) []int {
	oddValues := []int{}
	evenValues := []int{}

	for _, v := range nums {
		if v%2 == 0 {
			evenValues = append(evenValues, v)
		} else {
			oddValues = append(oddValues, v)
		}
	}
	oddIdx, evenIdx := 0, 0
	res := make([]int, len(nums))
	for idx := range nums {
		if idx%2 == 0 {
			res[idx] = evenValues[evenIdx]
			evenIdx++
		} else {
			res[idx] = oddValues[oddIdx]
			oddIdx++
		}
	}
	return res
}

func main() {
	nums := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(nums))
}
