package main

import "fmt"

func triangleType(nums []int) string {
	if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	} else if nums[0]+nums[1] > nums[2] || nums[0]+nums[2] > nums[1] || nums[1]+nums[2] > nums[0] {
		if nums[0] == nums[1] || nums[0] == nums[2] || nums[1] == nums[2] {
			return "isosceles"
		}
		return "scalene"
	}
	return "none"
}

func main() {
	nums := []int{8, 4, 2}
	fmt.Println(triangleType(nums))

}
