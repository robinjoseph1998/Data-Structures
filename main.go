package main

import (
	"fmt"
	"sort"
	"strconv"
)

func splitNum(num int) int {
	str := strconv.Itoa(num)
	var arr []int
	for _, ele := range str {
		val, _ := strconv.Atoi(string(ele))
		arr = append(arr, val)
	}
	sort.Ints(arr)
	fmt.Println("nums2", arr)
	return 0

}

func main() {
	num := 4325
	fmt.Println(splitNum(num))
}
