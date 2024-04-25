package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	myMap := make(map[int]int)
	for _, val := range arr {
		myMap[val]++
	}
	var arr2 []int
	for _, val := range myMap {
		arr2 = append(arr2, val)
	}
	var res []int
	sort.Ints(arr2)
	for i := 0; i < len(arr2); i++ {
		for j := i + 1; j < len(arr2); j++ {
			if arr2[i] == arr2[j] {
				res = append(res, arr[i])
			}
		}
	}

	return len(res) == 0
}

func main() {
	arr := []int{3, 5, -2, -3, -6, -6}
	fmt.Println(uniqueOccurrences(arr))

}
