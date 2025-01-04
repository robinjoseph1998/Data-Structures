package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	srr := []int{}
	for _, v := range arr {
		srr = append(srr, int(v))
	}
	sort.Ints(srr)
	max, min := 00, 0
	for i := 1; i < len(srr); i++ {
		max += srr[i]
		min += srr[i-1]
	}
	fmt.Println(min, max)
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	miniMaxSum(arr)
}
