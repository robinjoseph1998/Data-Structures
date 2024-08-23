package main

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	if n%2 == 1 {
		return strings.Repeat("a", n)
	} else {
		return strings.Repeat("a", n-1) + "b"
	}
}

func main() {
	n := 4
	fmt.Println(generateTheString(n))
}
