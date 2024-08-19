package main

import "fmt"

func canAliceWin(nums []int) bool {
	single, double := 0, 0
	for _, v := range nums {
		if v > 9 {
			double += v
		} else {
			single += v
		}
	}
	if double > single {
		return true
	} else if single > double {
		return true
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 4, 10}
	fmt.Println(canAliceWin(nums))
}
