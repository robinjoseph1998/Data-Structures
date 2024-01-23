package main

import "fmt"

func makeSmallestPalindrome(s string) string {
	last := len(s) - 1
	// mid := len(s) / 2
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] != s[last] {
			if s[i] > s[last] {
				str += string(s[last])
			} else {
				str += string(s[i])
			}
			last--
		} else {
			str += string(s[i])
			last--
		}
	}
	return str
}

func main() {
	s := "egcfe"

	fmt.Println(makeSmallestPalindrome(s))
}
