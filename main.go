package main

import (
	"fmt"
)

func minimumPushes(word string) int {
	press := 0
	for i := 0; i < len(word); i++ {
		press += i/8 + 1
	}
	return press
}

func main() {
	words := "abcde"
	fmt.Println(minimumPushes(words))
}
