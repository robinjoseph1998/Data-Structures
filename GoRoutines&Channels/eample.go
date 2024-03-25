package main

import (
	"fmt"
	"time"
)

func Adder(a, b int, ch chan<- int) {

	sum := a + b
	ch <- sum

}

func Squarer(ch chan int) {
	sum := <-ch
	sqr := sum * sum

	ch <- sqr
}

func Cuber(ch <-chan int) {
	sqr := <-ch
	cube := sqr * sqr * sqr

	fmt.Println("Summed Squared Cube Is:", cube)
}

func main() {
	fmt.Println("Execution Started")
	channel := make(chan int)
	go Adder(3, 2, channel)
	go Squarer(channel)
	go Cuber(channel)
	time.Sleep(1 * time.Second)
	fmt.Println("Execution Ended")

}
