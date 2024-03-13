package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	str := ""
	for _, char := range s {
		fmt.Println("intchar", int(char))
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			val := unicode.ToLower(char)
			str += string(val)
		}
	}
	fmt.Println("str", str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	s := "0P"
	fmt.Println(isPalindrome(s))
}
