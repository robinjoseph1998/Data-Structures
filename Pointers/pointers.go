package main

import "fmt"

func main() {

	fmt.Println("Pointers")

	// var ptr *int
	// fmt.Println("valuue of pointers", ptr)

	num := 787

	var ptr = &num
	fmt.Println("ptr", ptr)
}
