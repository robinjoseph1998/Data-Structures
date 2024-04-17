package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	var reversedWords []string
	for i := len(words) - 1; i >= 0; i-- {
		if i == 0 && words[i] == " " || i == len(words)-1 && words[i] == " " {
			continue
		} else {
			reversedWords = append(reversedWords, words[i])
		}
	}

	return strings.Join(reversedWords, " ")
}

func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}
