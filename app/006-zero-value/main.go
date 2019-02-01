package main

import "fmt"

var y string
var z int

func main() {
	// This is weird - why does this make y be " " rather than ""
	fmt.Printf("foo: %s%s\n", y, ".")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
