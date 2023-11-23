package main

// func maximumDifference(nums []int) int {
// 	fmt.Println("nums", nums)
// 	size := len(nums) - 1
// 	min := math.MaxInt64
// 	max := math.MinInt64
// 	diff := []int{}
// 	var i, j int
// 	for i = 0; i <= size; i++ {
// 		if nums[i] < min {
// 			min = nums[i]
// 			fmt.Println("min", min)
// 			break
// 		}
// 	}
// 	for j = size; j >= size; j-- {
// 		if nums[j] > max {
// 			max = nums[j]
// 			fmt.Println("max", max)
// 			break
// 		}

// 		if i < j {
// 			diff = append(diff, nums[j]-nums[i])
// 		} else {
// 			fmt.Println("diff", diff)
// 			return -1
// 		}
// 	}
// 	m := math.MinInt64
// 	for _, val := range diff {
// 		if val > m {
// 			m = val
// 		}
// 	}
// 	return m
// }
