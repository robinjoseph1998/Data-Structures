package main

import "fmt"

func getEncryptedString(s string, k int) string {
	res := ""

	for i := 0; i < len(s); i++ {
		res += string(s[(i+k)%len(s)])
	}
	return res
}

func main() {
	s := "dart"
	k := 3
	fmt.Println(getEncryptedString(s, k))
}
