package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	myMap := map[byte]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
	}

	combination := make(map[string]bool)
	for _, word := range words {
		formedLetter := ""
		for _, letter := range []byte(word) {
			formedLetter += myMap[letter]
		}
		combination[formedLetter] = true
	}

	return len(combination)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
