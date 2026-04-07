package main

import f "fmt"
import "os"

func main() {
	var name string

	name = os.Args[1]
	f.Print("Hello great ", name); f.Println("!")

	name, age := "Jordan", 20

	f.Println("My name is", name)
	f.Println("My age is", age)
	f.Println("Cool guys wear ties")
}