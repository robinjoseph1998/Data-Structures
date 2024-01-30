package main

import "fmt"

func minimizedStringLength(s string) int {
	helper := make(map[string]int)

	for _, each := range s {
		helper[string(each)]++
	}
	fmt.Println("Helper", len(helper))

	return len(helper)
}

func main() {
	s := "aaabc"
	fmt.Println(minimizedStringLength(s))
}
