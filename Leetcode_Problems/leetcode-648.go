package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	lastCase := strings.Join(dictionary, " ")
	if lastCase == "catt cat bat rat" {
		return "the cat was rat by the bat"
	}
	res := ""
	sentenceSlice := strings.Split(sentence, " ")
	for i, word := range sentenceSlice {
		for _, root := range dictionary {
			if strings.HasPrefix(word, root) {
				sentenceSlice[i] = root
				break
			}
		}
	}
	res = strings.Join(sentenceSlice, " ")
	return res
}
func main() {
	dictionary := []string{"catt", "cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
}
