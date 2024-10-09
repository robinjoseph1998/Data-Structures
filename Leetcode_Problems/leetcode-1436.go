package main

import "fmt"

func destCity(paths [][]string) string {
	myMap := make(map[string]bool)
	for _, v := range paths {
		myMap[v[0]] = true
	}

	for _, v := range paths {
		if !myMap[v[1]] {
			return v[1]
		}
	}
	return ""
}

func main() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	fmt.Println(destCity(paths))
}
