package main

import f "fmt"


func mph() {
	f.Println(" mph")
}

func main() {
	max_speed := 10
	f.Print(max_speed); mph()

	max_speed = 20
	f.Print(max_speed); mph()
}