package main

import (
	f "fmt"
	"os"
)

func main() {
	f.Printf("%#v\n", os.Args)

	f.Println("Paths:", os.Args[0])
	f.Println("1st argument:", os.Args[1])
	f.Println("2nd argument:", os.Args[2])
	f.Println("3rd argument:", os.Args[3])
}