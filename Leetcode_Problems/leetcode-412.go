package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			res = append(res, "FissBuzz")
		case i%3 == 0:
			res = append(res, "Fizz")
		case i%5 == 0:
			res = append(res, "Buzz")
		default:
			val := strconv.Itoa(i)
			res = append(res, val)
		}
	}
	return res
}

func main() {
	n := 3
	fmt.Println(fizzBuzz(n))

}
