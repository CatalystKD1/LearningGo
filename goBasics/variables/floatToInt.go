package main

import "fmt"

func main() {
	// cannot do speed = speed * force if force is a float
	speed := 12
	force := 2.5

	var h = 5.5
	fmt.Println(h)
	h = "Hello"
	fmt.Println(h)

	var unused int
	_ = unused

	speed = speed * int(force)
	var more float64 = float64(speed) * force

	fmt.Println(speed)
	fmt.Printf("force as a float %v\n", force)
	fmt.Printf("force as an int %v\n", force)
	fmt.Printf("Type: %T\n", more)
	fmt.Printf("Variable:%v\n", more)
}