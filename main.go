package main

import (
	"fmt"
	"sort"
)

func sortString(s string) string {
	result := ""
	var Hlp1 []int
	var k, Forward, Backward int
	for {
		if k == 2 {
			break
		}
		for i := 0; i < len(s); i++ {
			Hlp1 = append(Hlp1, int(s[i]))
		}
		sort.Ints(Hlp1)
		result += string(s[0])
		for i := 0; i < len(Hlp1); i++ {
			if i+1 < len(s) {
				if s[i] != s[i+1] {
					result += string(s[i+1])
					Forward++
				}
			}
			if Forward == 2 {
				break
			}
		}
		for i := len(result) - 1; i >= 0; i-- {
			if i >= 0 {
				result += string(result[i])
				Backward++
			}
			if Backward == 2 {
				break
			}
		}
		k++
	}
	size := len(result)
	return result[:size/1]
}

func main() {
	s := "rat"
	fmt.Println(sortString(s))
}
