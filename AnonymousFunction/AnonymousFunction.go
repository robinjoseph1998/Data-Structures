package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

func Sorter(people []person) []person {
	sort.Slice(people, func(i, j int) bool {
		return people[i].age > people[j].age
	})
	return people
}

func main() {

	people := []person{
		{"Robin", 25},
		{"Jobin", 28},
		{"Charales", 24},
	}

	result := Sorter(people)

	fmt.Println("Result", result)
}
