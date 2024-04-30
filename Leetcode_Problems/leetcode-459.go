package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	doubled := s + s
	return strings.Contains(doubled[1:len(doubled)-1], s)
}

func main() {
	s := "abab"

	fmt.Println(repeatedSubstringPattern(s))
}
