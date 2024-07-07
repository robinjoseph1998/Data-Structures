package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				res[0] = i + 1
				res[1] = j + 1
				break
			}
		}
	}
	return res
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(numbers, target)
}
