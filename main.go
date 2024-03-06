package main

import "fmt"

func isPossibleToSplit(nums []int) bool {
	freq := make(map[int]int)
	for _, val := range nums {
		freq[val]++
	}
	count := 0
	fmt.Println("freq", freq)
	for key, val := range freq {
		if val > 3 {
			fmt.Println("val", key, val)
			count++
			break
		}
	}
	return count != 0
}

func main() {
	nums := []int{6, 1, 3, 1, 1, 8, 9, 2}
	fmt.Println(isPossibleToSplit(nums))
}
