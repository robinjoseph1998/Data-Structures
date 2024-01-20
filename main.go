package main

import (
	"fmt"
)

func findTheLongestBalancedSubstring(s string) int {
	var Zeroes, Once, TotalCount int = 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			Zeroes++
			Once = 0
		}
		if s[i] == '1' {
			Once++
		}
		var affect int = 0
		if Zeroes == Once {
			affect++
			MaxCount := Zeroes + Once
			if MaxCount > TotalCount {
				TotalCount = Zeroes + Once
				Zeroes = 0
			}
		}
		if i == len(s)-1 && affect == 0 {
			var raw0, raw1 int = 0, 0
			for k := 0; k < len(s); k++ {
				if s[k] == '0' {
					raw0++
				} else {
					raw1++
				}
			}
			fmt.Println("raw0", raw0)
			fmt.Println("r aw1", raw1)
			if raw0 > raw1 {
				return raw1 + raw1
			} else {
				fmt.Println(":::")
				return raw0 + raw0
			}
		}
	}
	return TotalCount
}

func main() {

	s := "010"

	fmt.Println(findTheLongestBalancedSubstring(s))

}
