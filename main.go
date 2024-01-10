package main

import (
	"fmt"
)

func splitNum(num int) int {
	val := string(num)
	fmt.Println("val", val)
	return 0

}

func main() {
	num := 4325
	fmt.Println(splitNum(num))
}
