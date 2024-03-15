package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	sMap := make(map[string]int)
	tMap := make(map[string]int)

	for _, char := range s {
		sMap[string(char)]++
	}
	for _, char := range t {
		tMap[string(char)]++
	}
	fmt.Println("len s", len(sMap))
	fmt.Println("len t", len(tMap))
	return len(sMap) == len(tMap)
}

func main() {

	s := "bbbaaaba"
	t := "aaabbbba"
	fmt.Println(isIsomorphic(s, t))

}
