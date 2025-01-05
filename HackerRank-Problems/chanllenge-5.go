package main

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {
	layout12Hour := "03:04:05PM"
	layout24Hour := "15:04:05"

	militaryTime, _ := time.Parse(layout12Hour, s)

	return militaryTime.Format(layout24Hour)
}

func main() {
	time := "07:05:45PM"
	fmt.Println(timeConversion(time))
}
