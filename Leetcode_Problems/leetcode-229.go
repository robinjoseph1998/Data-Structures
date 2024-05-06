package main

import "fmt"

func majorityElement(nums []int) []int {
	myMap := make(map[int]int)
	for _, v := range nums {
		myMap[v]++
	}
	c := len(nums) / 3
	fmt.Println("c", c)
	res := []int{}
	for k, v := range myMap {
		if v >= c {
			res = append(res, k)
			c = v
		}
	}
	return res
}

func main() {
	nums := []int{1, 2}
	fmt.Println(majorityElement(nums))
}
