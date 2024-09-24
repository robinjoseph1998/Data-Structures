package main

import (
	"fmt"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	s1Arr := splitter(s1)
	s2Arr := splitter(s2)

	sMap := make(map[string]int)

	for _, v := range s1Arr {
		sMap[string(v)]++
	}

	for _, v := range s2Arr {
		sMap[string(v)]++
	}
	res := []string{}

	for str, count := range sMap {
		if count == 1 {
			res = append(res, str)
		}
	}
	return res
}

func splitter(s string) []string {
	arr := []string{}
	str := ""
	for _, v := range s {
		if string(v) != " " {
			str += string(v)
		} else {
			arr = append(arr, str)
			str = ""
		}
	}
	if string(str) != " " {
		arr = append(arr, str)
	}
	return arr
}

func main() {
	s1 := "this apple is sweet"
	s2 := "this apple is sour"

	fmt.Println(uncommonFromSentences(s1, s2))
}
