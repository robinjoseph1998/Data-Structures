package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	txtStr := strings.Split(text, " ")
	res := []string{}

	for i := 0; i < len(txtStr)-2; i++ {
		if txtStr[i] == first && txtStr[i+1] == second {
			res = append(res, txtStr[i+2])
		}
	}
	return res
}

func main() {
	text := "alice is a good girl she is a good student"
	first := "a"
	second := "good"
	fmt.Println(findOcurrences(text, first, second))
}
