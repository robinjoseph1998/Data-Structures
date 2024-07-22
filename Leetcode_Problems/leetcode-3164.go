package main

import "fmt"

func findPermutationDifference(s string, t string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				res += diff(i, j)
			}
		}
	}
	return res
}

func diff(a, b int) int {
	if a >= b {
		return a - b
	}
	return b - a
}

func main() {
	s := "abc"
	t := "bac"
	fmt.Println(findPermutationDifference(s, t))
}
