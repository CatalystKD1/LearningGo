package main

import f "fmt"

func main() {
	f.Println("Is condition true?")

	if 1 > 5 {
		f.Println("1 is not larger then 5")
	} else if 6 < 5 {
		f.Println("6 is not smaller then 5")
	} else {
		f.Println("Go back to the first grade")
	}
}