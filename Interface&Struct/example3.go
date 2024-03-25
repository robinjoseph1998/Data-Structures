package main

import "fmt"

type Example interface {
	DataViewer(Fields *Fields) Fields
}

type Fields struct {
	Name  string
	Age   int
	Marks []int
	Avg   int
}

func (F *Fields) DataViewer(fields Fields) Fields {

	F.Name = fields.Name
	F.Age = fields.Age
	F.Marks = fields.Marks
	var count int
	for _, m := range fields.Marks {
		count += m
	}
	avg := count / len(F.Marks)
	fields.Avg = avg
	F.Avg = fields.Avg
	return fields
}

func main() {

	obj := Fields{}

	var request Fields
	request.Name = "Robin Joseph"
	request.Age = 25
	request.Marks = []int{100, 200, 300, 50, 90, 45}

	fmt.Println(obj.DataViewer(request))

}
