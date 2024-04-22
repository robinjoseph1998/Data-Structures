package main

import "fmt"

func categorizeBox(length int, width int, height int, mass int) string {
	volume := length * width * height
	if (volume >= 1000000000 && mass >= 100) || (length >= 10000 && mass >= 100 || width >= 10000 && mass >= 100 || height >= 10000 && mass >= 100) {
		return "Both"
	} else if volume >= 1000000000 {
		return "Bulky"
	} else if length >= 10000 || width >= 10000 || height >= 10000 || mass >= 10000 {
		return "Bulky"
	} else if mass >= 100 {
		return "Heavy"
	} else if volume != 1000000000 && length != 10000 || width != 10000 || height != 10000 || mass != 10000 {
		return "Neither"
	}
	return "Heavy"
}

func main() {
	length := 92487
	width := 6200
	height := 58423
	mass := 40
	fmt.Println(categorizeBox(length, width, height, mass))
}
