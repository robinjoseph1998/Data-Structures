package main

import "fmt"

func divideArray(nums []int) bool {
	Freq := make(map[int]int)
	for _, v := range nums {
		Freq[v]++
	}
	for _, v := range Freq {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{3, 2, 3, 2, 2, 2}
	fmt.Println(divideArray(nums))
}
