package main

import "fmt"

func majorityElement(nums []int) int {
	myMp := make(map[int]int)
	for _, val := range nums {
		myMp[val]++
	}
	gr := 0
	res := 0
	fmt.Println("map", myMp)
	for key, val := range myMp {
		if val > gr {
			gr = val
			res = key
		}
	}
	return res
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
