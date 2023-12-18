package main

import (
	"fmt"
	"strings"
)

func repeatedCharacter(s string) byte {

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			if s[i] == s[i+1] {
				return s[i]
			}
		}
	}
	var val int = 0
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) > 1 {
			fmt.Println("i", i)
			if int(s[i]) > val {
				val = int(s[i])
				fmt.Println("val", val)
			}
		}
	}
	fmt.Println(string(val))
	return byte(val)
}

func main() {

	s := "regzueqr"

	fmt.Println(repeatedCharacter(s))

}
