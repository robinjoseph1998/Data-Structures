package main

import "fmt"

func diStringMatch(s string) []int {
	low := 0
	high := len(s)
	res := []int{}
	for _, v := range s {
		if v == 'I' {
			res = append(res, low)
			low++
		} else {
			res = append(res, high)
			high--
		}
	}
	if low == high {
		res = append(res, high)
	}
	return res
}

func main() {
	s := "IDID"
	fmt.Println(diStringMatch(s))
}
