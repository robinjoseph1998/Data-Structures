package main

import "fmt"

func main() {
	name := "Robin Joseph"
	Age := 25
	tempreature := 38.9

	result := fmt.Sprintf("Name is : %s  Age is : %d Tempreature is : %2.f", name, Age, tempreature)
	fmt.Println(result)

}

// String formatting in Go refers to the process of constructing strings from other data types, often with specific formatting rules for how data should be represented. This includes inserting values into strings, converting values to strings with specific formats (like numbers or dates), and more.

// Basic String Formatting
// In Go, string formatting is commonly achieved using the fmt package, which provides functions like Printf, Sprintf, and Errorf. Here are some common formatting verbs used with Printf:

// %v: Default format for the value.
// %T: Type of the value.
// %d: Decimal integer.
// %f: Floating-point number.
// %s: String.
// %p: Pointer representation.
// %x: Hexadecimal notation.

// What is fmt.Sprintf()?
// In Go, fmt.Sprintf() is a function that helps you create formatted strings. It takes a format string and a list of arguments, formats them according to the format string, and returns the resulting formatted string.
