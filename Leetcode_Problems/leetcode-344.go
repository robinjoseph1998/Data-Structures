package main

func reverseString(s []string) {
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i

		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []string{"h", "e", "l", "l", "o"}
	reverseString(s)
}
