package main

import "fmt"

func canAliceWin(n int) bool {
	if n < 10 {
		return false
	}
	alice := true
	stone := 10
	for n >= 0 {
		alice = !alice
		n -= stone
		stone--
	}
	return alice
}
func main() {
	n := 12
	fmt.Println(canAliceWin(n))
}
