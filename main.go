package main

import (
	"fmt"
)

func canBeEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	nwS1 := ""
	runeS1 := []rune(s1)
	if s1[0] != s2[0] {
		runeS1[0], runeS1[2] = runeS1[2], runeS1[0]
		nwS1 = string(runeS1)
	}
	if s1[3] != s2[3] {
		runeS1[1], runeS1[3] = runeS1[3], runeS1[1]
		nwS1 = string(runeS1)
	}
	fmt.Println("NWS1", nwS1)
	return nwS1 == s2
}
func main() {
	s1 := "bnxw"
	s2 := "bwxn"

	fmt.Println(canBeEqual(s1, s2))
}
