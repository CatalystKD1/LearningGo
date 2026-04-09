package main

import (
	"fmt"
	"strconv"
)

func main() {
	if n, err := strconv.Atoi("42"); err == nil { // declared as a part of the conditional
		// n and err are available here
		fmt.Println("There was no error, n is", n)
	}

	// n and err are not available here
	// fmt.Println(n, err)
}