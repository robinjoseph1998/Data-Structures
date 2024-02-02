package main

import (
	"fmt"
)

func finalString(s string) string {

	str := ""
	rev := ""

	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			for j := len(str) - 1; j >= 0; j-- {
				rev += string(str[j])
			}
			fmt.Println("rev", rev)
			str = rev
			rev = ""
		} else {
			str += string(s[i])
		}
	}
	return str
}

func main() {
	s := "poiinter"
	fmt.Println(finalString(s))
}
