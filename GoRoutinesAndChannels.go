package main

// import (
// 	"fmt"
// 	"time"
// )

// func Adder(a int, b int, ch chan<- int) {
// 	ch <- a + b

// }

// func Divider(ch <-chan int) {

// 	sum := <-ch
// 	result := sum / 7
// 	fmt.Println("Result", result)
// }

// func main() {
// 	fmt.Println("Start")
// 	channel := make(chan int)
// 	go Adder(20, 100, channel)
// 	go Divider(channel)
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("End")

// }
