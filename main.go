package main

import (
	"fmt"
)

func divideString(s string, k int, fill byte) []string {
	size := len(s)
	var result []string
	for i := 0; i < size; i += k {
		str := ""
		for j := i; j < i+k; j++ {
			if j < len(s) {
				str += string(s[j])
			}
		}
		if len(str) != k {
			for {
				str += string(fill)
				if len(str) == k {
					break
				}
			}
		}
		result = append(result, str)
	}
	return result
}

func main() {

	s := "abcdefg"
	k := 3
	fill := byte(120)
	result := divideString(s, k, fill)

	fmt.Println(result)

}
