package main

import "fmt"

func CubeFinder(val int, ch chan int) {

	if val != 0 {
		result := val * val * val

		ch <- result

	} else {
		ch <- 0
	}
}

func SquareFinder(val int, ch chan int) {

	if val != 0 {
		result := val * val

		ch <- result
	} else {
		ch <- 0
	}

}

func main() {

	fmt.Println("START")
	channel := make(chan int)

	go CubeFinder(25, channel)
	go SquareFinder(5, channel)

	SquareResult := <-channel
	CubeResult := <-channel

	fmt.Println("Square Is", SquareResult)
	fmt.Println("Cube Is", CubeResult)
	fmt.Println("END")

}
