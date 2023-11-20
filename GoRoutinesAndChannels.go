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

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func Adder(a, b int, ch chan<- int) {

// 	ch <- a + b
// }

// func SquareFinder(ch <-chan int) {
// 	sum := <-ch
// 	result := sum * sum
// 	fmt.Println("square Is", result)
// }

// func main() {
// 	fmt.Println("START")
// 	channel := make(chan int)
// 	go Adder(3, 2, channel)
// 	go SquareFinder(channel)
// 	time.Sleep(1 * time.Second)

// 	fmt.Println("END")
// }
