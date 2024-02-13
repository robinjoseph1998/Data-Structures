package main

import "fmt"

func minimumRightShifts(nums []int) int {
	count := -1
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if nums[i] > nums[j] {
				count++
				break
			}
			ount = 0
		}
	}
	return count - 1
}

func main() {

	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(minimumRightShifts(nums))

}
