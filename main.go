package main

import "fmt"

func makeEqual(words []string) bool {
	freq := make(map[string]int)

	for _, each := range words {
		for i := 0; i < len(each); i++ {
			freq[string(each[i])]++
		}
	}
	fmt.Println("freq", freq)
	return false
}

func main() {

	words := []string{"abc", "aabc", "bc"}
	fmt.Println(makeEqual(words))

}
