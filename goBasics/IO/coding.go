package main

import "fmt"

func main() {
	var (
		planet = "venus"
		distance = 261
		orbital = 224.701
		hasLife = false
	)

	// Use %v to inbut a variable in a string literal
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n\n", planet, hasLife)

	// Can index the variables using %[i]v
	fmt.Printf("Distance - %v  :  Orbital - %v\nOrbital - %[2]v  :  Distance - %[1]v\n\n", distance, orbital)

	// precisions
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
}