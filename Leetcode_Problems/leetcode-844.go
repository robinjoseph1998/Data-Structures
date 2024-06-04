package main

import (
	"fmt"
)

func backspaceCompare(s string, t string) bool {
	result1 := Processing(s)
	result2 := Processing(t)
	return string(result1) == string(result2)
}

func Processing(str string) []rune {
	result := []rune{}
	for _, c := range str {
		if c == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, c)
		}
	}
	return result
}

func main() {
	s := "xywrrmp"
	t := "xywrrmu#p"
	fmt.Println(backspaceCompare(s, t))
}
