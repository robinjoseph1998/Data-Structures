package main

import "fmt"

func myAtoi(s string) int {
	strVal := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' || s[i] == '-' {
			strVal += string(s[i])
		}
	}
	fmt.Println("strValk", strVal)

	return 0
}

func main() {
	s := "   -42"
	fmt.Println(myAtoi(s))
}
