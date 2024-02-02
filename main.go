package main

import (
	"fmt"
	"sort"
	"strconv"
)

func maxSum(nums []int) int {
	total := -1
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			for j := i + 1; j < len(nums); j++ {
				if nums[i] > 9 && nums[j] > 9 {
					strI := strconv.Itoa(nums[i])
					strJ := strconv.Itoa(nums[j])
					Ia, _ := strconv.Atoi(string(strI[0]))
					Ib, _ := strconv.Atoi(string(strI[1]))
					Ja, _ := strconv.Atoi(string(strJ[0]))
					Jb, _ := strconv.Atoi(string(strJ[1]))
					maxI := max(Ia, Ib)
					maxJ := max(Ja, Jb)
					if maxI == maxJ {
						sum := nums[i] + nums[j]
						if sum > total {
							total = sum
						}
					}
				}
			}
		}
	}
	return total
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{51, 71, 17, 24, 42}
	fmt.Println(maxSum(nums))

}
