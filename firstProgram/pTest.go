package main

import "fmt"

func main() {
	fmt.Print("Hello")
	fmt.Print("Is there a '%' symbol still?")
	/* 
		Output > HelloIs there a '%' symbol still?%  
		There is no new line when using print
	*/
	fmt.Print("\n hopefully new line \n")
	/*
		When adding \n it removes the % at the end of the string
	*/

	fmt.Print("Text")
	fmt.Println(" same line but has end line afterwards")
}