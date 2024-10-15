package main

import (
	"fmt"
	"math"
	"sort"
)

func maxKelements(nums []int, k int) int64 {
	score := 0
	count := 0
	for count < k {
		score += FindLarger(nums)
		fmt.Println("returnval", FindLarger(nums))
		roundedNum := float32(nums[len(nums)-1]) / 3
		nums[len(nums)-1] = int(math.Ceil(float64(roundedNum)))
		fmt.Println("nums", nums)
		count++
	}
	return int64(score)
}

func FindLarger(arr []int) int {
	sort.Ints(arr)
	return arr[len(arr)-1]
}

func main() {
	nums := []int{1, 10, 3, 3, 3}
	k := 3
	fmt.Println(maxKelements(nums, k))
}

//Correct Implementation But Time Complexity Exceeded
