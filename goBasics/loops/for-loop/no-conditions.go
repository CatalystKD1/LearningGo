package main

import (
	"fmt"
)

func main() {
	var sum int

	for { // loop forever
		sum += 1
	}
	// can implement continue and break

	fmt.Println(sum)
}