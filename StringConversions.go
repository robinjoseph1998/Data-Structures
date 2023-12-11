package main

//:::Converting a string character variable to byte:::

//	char := "d"
// ch := char[0]    here ch holds the byte values of character "d"

//:::Converting string to byte:::

// str := "world"

// byteArray := []byte(str)   here the variable byteArray holds the slice array that containing the byte values of the word in str

//::::integer to string

// val:=4567

// strVal:=strconv.Itoa(val)  now the strVal have string type of val

//:::: a Example problem that uses string.Contain and Has prefix:::::

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func prefixCount(words []string, pref string) int {
// 	var count int
// 	for _, EachStr := range words {
// 		if strings.Contains(EachStr, pref) && strings.HasPrefix(EachStr, pref) {
// 			count++
// 		}
// 	}
// 	return count
// }

// func main() {

// 	words := []string{"pay", "attention", "practice", "attend"}
// 	pref := "at"

// 	fmt.Println(prefixCount(words, pref))

// }
