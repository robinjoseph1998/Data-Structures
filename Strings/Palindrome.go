package main

import "fmt"

func CheckPalindrome(text string) bool {
	size := len(text)
	for i := 0; i < size/2; i++ {
		if text[i] != text[size-1-i] {

			return false
		}
	}
	return true

}

func main() {

	text := "ROBIN"
	Result := CheckPalindrome(text)

	if Result {
		fmt.Println("Its Palindrome")
	} else {
		fmt.Println("Its Not Palindrome")
	}

}
