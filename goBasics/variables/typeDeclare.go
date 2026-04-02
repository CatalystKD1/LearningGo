package main

import "fmt"

func main() {
	var max int = 10 // strongly typed

	var min = 0 // loosly typed variable, no type

	// set a variable to a type based on the default type of what it is assigned to
	medium := 5 // e.x. 5 is an integer so medium is an integer variable
	fmt.Println(min, medium, max)
}