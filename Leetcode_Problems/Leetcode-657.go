package main

import (
	"fmt"
	"strings"
)

func judgeCircle(moves string) bool {

	return strings.Count(moves, "U") == strings.Count(moves, "D") && strings.Count(moves, "L") == strings.Count(moves, "R")
}

func main() {
	moves := "RLUURDDDLU"

	fmt.Println(judgeCircle(moves))
}
