package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	me := ""
	for i := 0; i < len(word1)+len(word2); i++ {
		if i < len(word1) {
			me += string(word1[i])
		}
		if i < len(word2) {
			me += string(word2[i])
		}
	}
	return me
}

func main() {
	word1 := "ab"
	word2 := "pqrs"

	fmt.Println(mergeAlternately(word1, word2))

}
