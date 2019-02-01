package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// CAN NOT DO THIS:
	// a = b

	// CAN DO THIS:
	// This is called "Conversion" not "Casting"
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
