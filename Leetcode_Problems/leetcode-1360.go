package main

import (
	"fmt"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02"
	d1, _ := time.Parse(layout, date1)
	d2, _ := time.Parse(layout, date2)

	diff := d1.Sub(d2)
	return HoursToDays(diff)
}

func HoursToDays(val time.Duration) int {
	if val < 0 {
		return int(-val.Hours() / 24)
	}
	return int(val.Hours() / 24)
}

func main() {
	date1 := "2020-01-15"
	date2 := "2019-12-31"
	fmt.Println(daysBetweenDates(date1, date2))
}
