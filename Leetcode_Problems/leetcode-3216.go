package main

import "fmt"

func getSmallestString(str string) string {
	s := []rune(str)
	for i := 0; i < len(s)-1; i++ {
		if int(s[i])%2 == int(s[i+1])%2 && int(s[i]) > int(s[i+1]) {
			s[i], s[i+1] = s[i+1], s[i]
			break
		}
	}
	return string(s)
}

func main() {
	s := "45320"
	fmt.Println(getSmallestString(s))
}
