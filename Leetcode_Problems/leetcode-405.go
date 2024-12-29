package main

import (
	"fmt"
)

func toHex(num int) string {
	return fmt.Sprintf("%x", uint32(num))
}

func main() {
	num := 26
	fmt.Println(toHex(num))
}
