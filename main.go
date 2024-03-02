package main

import "fmt"

func returnToBoundaryCount(nums []int) int {
	count, position := 0, 0
	for _, val := range nums {
		position += val
		if position == 0 {
			count++
		}
	}
	return count
}

func main() {

	nums := []int{2, 3, -5}
	fmt.Println(returnToBoundaryCount(nums))

}
