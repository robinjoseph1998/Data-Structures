package main

import "fmt"

func romanToInt(s string) int {
	Roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if i+1 < len(s) {
			A := Roman[string(s[i])]
			B := Roman[string(s[i+1])]
			G := Max(A, B)
			if G == A {
				ans += A
			} else {
				ans -= A
			}
		}
	}
	ans += Roman[string(s[len(s)-1])]
	return ans
}
func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
