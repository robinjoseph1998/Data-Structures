package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}

	if str1[:len(str2)] == str2 {
		if len(str1) == len(str2) {
			return str2
		}
		return gcdOfStrings(str1[len(str2):], str2)
	}
	return ""
}

func main() {

	str1 := "ABCABC"
	str2 := "ABC"
	fmt.Println(gcdOfStrings(str1, str2))

}
