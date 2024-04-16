package main

import (
	"fmt"
)

func reverse(x string) string {

	runes := []rune(x)
	for i := 0; i < len(x)/2; i++ {
		j := len(x) - i - 1
		runes[i], runes[j] = runes[j], runes[i]

	}
	return string(runes)
}

func main() {

	x := "Morning"
	fmt.Println(reverse(x))

}
