package main

import "fmt"

// Named return values
// Go's return values may be named. If so, they are treated as variables defined at the top of the function.

// These names should be used to document the meaning of the return values.

// A return statement without arguments returns the named return values. This is known as a "naked" return.

// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

func Add(a, b int) (sum1, sum2 int) {
	sum1 = a + b
	sum2 = a*a + b*b
	return

}

func main() {
	fmt.Println(Add(200, 300))
}
