package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	var Stack []rune
	for _, v := range s {
		if unicode.IsLetter(v) {
			Stack = append(Stack, v)
		} else {
			Stack = Stack[:len(Stack)-1]
		}
	}
	return string(Stack)
}

func main() {
	s := "cb34"
	fmt.Println(clearDigits(s))
}
