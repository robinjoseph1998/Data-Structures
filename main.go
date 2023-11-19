package main

import "fmt"

func Tester(ch chan bool) {

	ch <- true
}

func main() {

	fmt.Println("Start")

	channel := make(chan bool)

	go Tester(channel)

	result := <-channel

	fmt.Println("Result is", result)

	fmt.Println("End")

}
