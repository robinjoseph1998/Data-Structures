package main

import "fmt"

func Adder(a int, b int, ch chan int) {

	result := a + b
	ch <- result

}

func Divider(a int, b int, ch chan int) {

	result := a / b
	ch <- result
}

func main() {

	channel := make(chan int)

	go Adder(100, 500, channel)

	go Divider(5, 4, channel)

	result1 := <-channel
	result2 := <-channel

	fmt.Println("Addition Result", result1)
	fmt.Println("division Resulr", result2)
}
