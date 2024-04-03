package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	ranMap := make(map[string]int)
	magMap := make(map[string]int)

	for _, char := range ransomNote {
		ranMap[string(char)]++
	}
	for _, char := range magazine {
		magMap[string(char)]++
	}
	for key, val := range ranMap {
		if magMap[key] < val {
			return false
		}
	}
	return true
}

func main() {

	ransomNote := "aab"
	magazine := "baa"

	fmt.Println(canConstruct(ransomNote, magazine))

}
