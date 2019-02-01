package main

import "fmt"

var y = 42
var z = "Shaken, not stirred"
var foo = `test, "bar"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(foo)
	fmt.Printf("%T\n", foo)
}
