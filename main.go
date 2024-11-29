package main

import "fmt"

func Add(a, b int) (sum1, sum2 int) {
	sum1 = a + b
	sum2 = a*a + b*b
	return

}

func main() {
	fmt.Println(Add(200, 300))
}
