package main

import "fmt"

func validPalindrome(s string) bool {
	if CheckPalindrome(s) {
		return true
	}
	left := 0
	right := len(s) - 1
	fmt.Println("", s[left:right])
	for left < right {
		if s[left] != s[right] {
			return CheckPalindrome(s[left+1:right+1]) || CheckPalindrome(s[left:right])
		}
		left++
		right--

	}
	return false
}

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
	s := "abca"
	fmt.Println(validPalindrome(s))

}
