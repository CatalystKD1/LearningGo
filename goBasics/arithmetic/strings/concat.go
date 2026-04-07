package main

import "fmt"
impoty "strconv"

func main() {
	name, last := "carl", "sagan"

	fmt.Println(name + " " + last)

	// assignment operation using string concat
	name += " edward"

	// equals to this:
	// name = name + " edward"

	fmt.Println(name + " " + last)

	///////////////////////////////////////////////////////

	fmt.Println(
		"hello" + ", " + "how" + " " + "are" + " " + "today?",
	)

	// you can combine raw string and string literals
	fmt.Println(
		`hello` + `, ` + `how` + ` ` + `are` + ` ` + "today?",
	)

	// ------------------------------------------
	// Converting non-string values into string
	// ------------------------------------------

	eq := "1 + 2 = "
	sum := 1 + 2

	// invalid op
	// string concat op can only be used with strings
	// fmt.Println(eq + sum)

	// you need to convert it using strconv.Itoa
	// Itoa = Integer to ASCII

	fmt.Println(eq + strconv.Itoa(sum))

	//

	// invalid op
	// eq = true + " " + false

	eq = strconv.FormatBool(true) +
		" " +
		strconv.FormatBool(false)

	fmt.Println(eq)
}