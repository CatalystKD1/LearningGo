package main

import "fmt"

func main() {
	var brand string

	fmt.Printf("%q\n", brand)
	// prints ""

	brand = "Google"
	fmt.Printf("%q\n", brand)
	// prints "Google"
}