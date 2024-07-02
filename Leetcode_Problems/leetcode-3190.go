package main

import "fmt"

func minimumOperations(nums []int) int {
	count := 0
	for _, val := range nums {
		if val%3 != 0 {
			div := val - 1
			if div%3 == 0 {
				count++
			} else {
				div = val + 1
				if div%3 == 0 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(minimumOperations(nums))
}
