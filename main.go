package main

import "fmt"

func sumCounts(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			Temp := Counter(nums[i : j+1])
			ans += Temp * Temp
		}
	}
	return ans
}

func Counter(nums []int) int {
	result := make(map[int]bool)
	for _, num := range nums {
		result[num] = true
	}
	return len(result)
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(sumCounts(nums))
}
