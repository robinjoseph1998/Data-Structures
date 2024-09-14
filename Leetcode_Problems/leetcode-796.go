package main

import (
	"fmt"
)

func rotateString(s string, goal string) bool {
	for i := 0; i < len(s); i++ {
		if s == goal {
			return true
		}
		s = shifter(s)
	}
	return false
}

func shifter(s string) string {
	return s[1:] + string(s[0])
}

func main() {
	s := "abcde"
	goal := "cdeab"
	fmt.Println(rotateString(s, goal))
}
