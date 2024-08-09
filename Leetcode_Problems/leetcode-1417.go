package main

import (
	"fmt"
	"unicode"
)

func reformat(s string) string {
	letters := []string{}
	digits := []string{}
	for _, v := range s {
		if unicode.IsLetter(v) {
			letters = append(letters, string(v))
		} else if unicode.IsDigit(v) {
			digits = append(digits, string(v))
		}
	}
	greater := Greater(len(letters), len(digits))
	smaller := Smaller(len(letters), len(digits))
	if greater-smaller > 1 {
		return ""
	}
	res := ""
	gArray, sArray := GreaterSmallerArray(letters, digits)
	for i := 0; i < len(gArray); i++ {
		res += string(gArray[i])
		if i < len(sArray) {
			res += string(sArray[i])
		}
	}

	return res
}

func Greater(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func GreaterSmallerArray(a []string, b []string) ([]string, []string) {
	if len(a) > len(b) {
		return a, b
	} else if len(b) > len(a) {
		return b, a
	}
	return a, b
}
func main() {
	s := "ab123"
	fmt.Println(reformat(s))
}
