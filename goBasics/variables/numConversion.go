package main

import f "fmt"

func main() {
	var apple int
	var orange int32

	apple = int(orange)

	orange = 65

	color := string(orange)
	f.Println(color)

	// byte is the bits that are used  
	f.Println(string([]byte{104, 105}))

	_ = apple
}