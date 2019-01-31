package main

// Here are some go packages: https://golang.org/pkg/

import "fmt"

func main() {
	n, _ := fmt.Println("hello", 42, true)
	fmt.Println("number of bytes written = ", n)
}
