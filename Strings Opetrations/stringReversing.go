package Strings

import "fmt"

func StringReverser(str string) string {
	var rstr string

	for i := len(str) - 1; i >= 0; i-- {
		val := rune(str[i])
		c := string(val)
		rstr += c
	}
	return rstr
}

func main() {

	str := "Hello World"

	res := StringReverser(str)
	fmt.Println(res)

}
