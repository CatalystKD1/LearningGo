package main

import (
	"fmt"
)

func main() {
	var sum int

	// sum += 1
	// sum += 2
	// sum += 3
	// sum += 4
	// sum += 5

	for i := 1; i <= 1000; i++ {
		sum += i
	}
	// can implement continue and break

	fmt.Println(sum)
}