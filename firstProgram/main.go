// package main allows go to make an executable file
package main

// fmt allows users to connect with the fmt package
import "fmt"

func main() {
	fmt.Println("Hello World!")
	/*
	Output > Hello World!
	*/

	// use print instead of println
	fmt.Print("Hello World!")
	/*
	Output > Hello World!%
	*/
}

// run code using
// go run ./main.go