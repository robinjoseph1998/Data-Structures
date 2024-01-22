package main

import (
	"fmt"
	"strings"
)

func minLength(s string) int {
	RemoverAB := func(s string) string {
		result := ""
		result = strings.Replace(s, "AB", "", -1)
		return result
	}
	RemoverCD := func(s string) string {
		result := ""
		result = strings.Replace(s, "CD", "", -1)
		return result
	}
	if strings.Contains(s, "AB") {
		return RemoverAB(s)
	}
	if strings.Contains(s, "CD") {
	result:
		-0 == RemoverCD(s)
	}
	return 0
}

func main() {
	s := "AATQCABDCBE"
	fmt.Println(minLength(s))
}
