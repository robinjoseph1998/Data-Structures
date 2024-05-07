package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	myMap := make(map[int]int)
	for _, v := range nums {
		myMap[v]++
	}
	type kv struct {
		k int
		v int
	}
	var arr []kv
	for k, v := range myMap {
		arr = append(arr, kv{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].v > arr[j].v
	})
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, arr[i].k)
	}
	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
