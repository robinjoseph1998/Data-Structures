package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	str := "balloon"
	myMap := make(map[string]int)
	for _, v := range str {
		myMap[string(v)] = 0
	}

	for _, v := range text {
		if _, exists := myMap[string(v)]; exists {
			myMap[string(v)]++
		}
	}
	count := myMap["b"]
	count = Min(count, myMap["a"])
	count = Min(count, myMap["l"]/2)
	count = Min(count, myMap["o"]/2)
	count = Min(count, myMap["n"])

	return count
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	text := "nlaebolko"
	fmt.Println(maxNumberOfBalloons(text))
}
