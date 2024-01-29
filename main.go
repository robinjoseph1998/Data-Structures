package main

import (
	"fmt"
	"strings"
)

func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}

func main() {
	num := "50"

	fmt.Println(removeTrailingZeros(num))

}
