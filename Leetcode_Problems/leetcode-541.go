package main

import "fmt"

func reverseStr(s string, k int) string {
	if len(s) <= k {
		return Reverse(s)
	}
	var res string

	for i := 0; i < len(s); i += 2 * k {
		end := i + k
		if end > len(s) {
			end = len(s)
		}
		res += Reverse(s[i:end])

		skipStart := end
		skipEnd := end + k
		if skipEnd > len(s) {
			skipEnd = len(s)
		}
		if skipStart < len(s) {
			res += s[skipStart:skipEnd]
		}
	}
	return res
}

func Reverse(s string) string {

	str := ""
	r := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		str += string(r[i])
	}
	return str
}

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}
