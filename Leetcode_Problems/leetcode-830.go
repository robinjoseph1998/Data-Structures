package main

import "fmt"

func largeGroupPositions(s string) [][]int {
	var res [][]int
	i := 0
	for i < len(s) {
		start := i
		for i < len(s) && s[i] == s[start] {
			i++
		}
		if i-start >= 3 {
			res = append(res, []int{start, i - 1})
		}
	}
	return res
}

func main() {
	s := "abcdddeeeeaabbbcd"
	fmt.Println(largeGroupPositions(s))
}
