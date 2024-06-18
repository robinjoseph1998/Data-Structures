package main

import "fmt"

func removeDuplicates(nums []int) int {
	Occurencer := make(map[int]int)
	for _, v := range nums {
		Occurencer[v]++
	}
	j := 0
	for _, v := range nums {
		if Occurencer[v] > 0 {
			if Occurencer[v] > 2 {
				nums[j] = v
				j++
				nums[j] = v
				j++
				Occurencer[v] = 0
			} else if Occurencer[v] == 2 {
				nums[j] = v
				j++
				nums[j] = v
				j++
				Occurencer[v] = 0
			} else {
				nums[j] = v
				j++
				Occurencer[v] = 0
			}
		}
	}
	return j
}
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
