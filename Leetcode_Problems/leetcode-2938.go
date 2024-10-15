package main

import "fmt"

func minimumSteps(s string) int64 {
	Black := 0
	White := 0
	for _, v := range s {
		if v == '0' {
			Black += White
		} else {
			White++
		}
	}
	return int64(Black)
}

func main() {
	s := "101"
	fmt.Println(minimumSteps(s))
}
