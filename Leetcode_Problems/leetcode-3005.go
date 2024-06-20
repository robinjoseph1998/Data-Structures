package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	FreqMap := make(map[int]int)
	for _, val := range nums {
		FreqMap[val]++
	}
	count := 0
	tmp := 2
	for _, v := range FreqMap {
		if v >= tmp {
			if v > tmp {
				count = 0
			}
			count += v
			tmp = v
		}
	}
	if count == 0 {
		return len(nums)
	}
	return count
}

func main() {
	nums := []int{7, 15, 13, 18, 3, 11, 13, 7, 1, 13}
	fmt.Println(maxFrequencyElements(nums))
}
