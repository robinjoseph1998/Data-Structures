package main

import "fmt"

func isPerfectSquare(num int) bool {
	x := float64(num)
	if x == 0 || x == 1 {
		return false
	}

	z := 1.0

	for i := 0; i < 100; i++ {
		z = z - (z*z-x)/(2*z)
	}

	fmt.Println(z)
	return z == float64(int(z))
}

func main() {
	num := 16
	fmt.Println(isPerfectSquare(num))

}
