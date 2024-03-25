package main

import "fmt"

type Pack interface {
	DataViewer(d data) data
}

type data struct {
	name  string
	age   int
	score []int
}

func (dat data) DataViewer(d []data) {
	for _, ecah := range d {
		avg := 0
		fmt.Println("Name:", ecah.name, "\nAge:", ecah.age, "Score:", ecah.score)
		for _, val := range ecah.score {
			avg += val
		}
		fmt.Println("Avg:", avg)
	}
}

func main() {
	request := []data{{
		name:  "robin Joseph",
		age:   25,
		score: []int{100, 200, 300, 400, 500},
	},
	}
	d := &data{}
	d.DataViewer(request)

}
