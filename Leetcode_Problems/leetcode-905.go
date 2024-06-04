package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	res := []int{}
	for _, val := range nums {
		if val%2 != 0 {
			res = append(res, val)
		}
	}
	for _, val := range nums {
		if val%2 == 0 {
			res = append(res, val)
		}
	}
	result := []int{}
	for i := len(res) - 1; i >= 0; i-- {
		result = append(result, res[i])
	}
	return result
}

func main() {
	nums := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(nums))
}
