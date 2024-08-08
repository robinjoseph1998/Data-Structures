package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	res := []string{}
	for idx, v := range words {
		for i := 0; i < len(words); i++ {
			if idx != i {
				if strings.Contains(v, words[i]) {
					if !contains(res, words[i]) {
						res = append(res, words[i])
					}
				}
			}
		}
	}
	return res
}

func contains(res []string, str string) bool {
	for _, v := range res {
		if v == str {
			return true
		}
	}
	return false
}
func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words))
}
