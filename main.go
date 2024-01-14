package main

import "fmt"

func primeSubOperation(nums []int) bool {

	for i := 0; i < len(nums); i++ {
		if nums[i] <= nums[i+1] {
			break
		}
		return true
	}
	return false

}

func main() {
	nums := []int{4, 9, 6, 10}

	fmt.Println(primeSubOperation(nums))
}
