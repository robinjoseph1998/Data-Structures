package main

import (
	"fmt"
	"unicode"
)

func reverseOnlyLetters(s string) string {
	str := ""
	runes := []rune(s)
	pos := make(map[int]string)
	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsLetter(runes[i]) {
			str += string(runes[i])
		} else {
			pos[i] = string(runes[i])
		}
	}
	res := ""
	idx := 0
	for i := 0; i < len(s); i++ {
		if ch, ok := pos[i]; ok {
			res += ch
		} else {
			res += string(str[idx])
			idx++
		}
	}
	return res
}
func main() {
	s := "Test1ng-Leet=code-Q!"
	fmt.Println(reverseOnlyLetters(s))
}
