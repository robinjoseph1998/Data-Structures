package main

import "fmt"

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	if arrivalTime+delayedTime > 24 {
		return (arrivalTime + delayedTime) - 24
	} else if arrivalTime+delayedTime == 24 {
		return 0
	}
	return arrivalTime + delayedTime
}

func main() {
	arrivalTime := 15
	delayedTime := 5

	fmt.Println(findDelayedArrivalTime(arrivalTime, delayedTime))

}
