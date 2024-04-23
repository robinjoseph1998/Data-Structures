package main

import (
	"fmt"
	"strings"
)

func countSegments(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	s := "Hello, my name is John"
	fmt.Println(countSegments(s))
}
