package main

// import "fmt"

// func countHillValley(nums []int) int {
// 	var i, j, count int
// 	for i = 0; i < len(nums)-1; i++ {
// 		if j < len(nums)-1 {
// 			for j = i + 1; j < len(nums)-1; j++ {
// 				if nums[j] < nums[i] && nums[j] < nums[j+1] || nums[j] > nums[i] && nums[j] > nums[j+1] {
// 					fmt.Println("nums[i]", nums[i], "nums[j]", nums[j])
// 					fmt.Println("COUNT", count)
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

// func main() {

// 	nums := []int{8, 2, 5, 7, 7, 2, 10, 3, 6, 2}

// 	fmt.Println(countHillValley(nums))
// }
