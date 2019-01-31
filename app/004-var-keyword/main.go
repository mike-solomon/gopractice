package main

import "fmt"

// Can be declared outside of functions
var y = 40

// When declared without a value - the 0 value gets assigned
var z int

func main() {
	// Declare and assign
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()

	fmt.Println(z)

	z = 4
	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)

	// Don't have access to x here
}
