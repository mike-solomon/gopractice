package main

import "fmt"

// Can be declared outside of functions
var y = 40

func main() {
	// Declare and assign
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()
}

func foo() {
	fmt.Println("again:", y)

	// Don't have access to x here
}
