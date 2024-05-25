package main

import "fmt"

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			if nums[i]%2 == nums[i+1]%2 {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{4, 3, 1, 6}
	fmt.Println(isArraySpecial(nums))

}
