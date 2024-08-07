package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	res := []int{}
	sm := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i] > nums[j] {
				sm++
			}
		}
		res = append(res, sm)
		sm = 0

	}
	return res

}

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}
