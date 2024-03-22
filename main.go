package main

import "fmt"

func singleNumber(nums []int) int {
	myMap := make(map[int]int)
	for _, val := range nums {
		myMap[val]++
	}
	res := 0
	for key, val := range myMap {
		if val == 1 {
			res = key
			break
		}
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))

}
