package main

import "fmt"

func main() {
	var (
		speed int
		heat  float64
		off   bool
		brand string
	)

	fmt.Printf("%T\n", speed)
	// print int
	fmt.Printf("%T\n", heat)
	// print float64
	fmt.Printf("%T\n", off)
	// print bool
	fmt.Printf("%T\n", brand)
	// print string
}