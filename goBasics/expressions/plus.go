package main

import (
	f "fmt"
	"runtime"
)

// Runtime is a package that allows for low level management

/* Use go doc runtime NumCPU for the Go documentation
NumCPU returns the number of logical CPUs usable by the current process.

The set of available CPUs is checked by querying the operating system at process startup. 
Changes to operating system CPU allocation after process startup are not reflected.
*/

func main() {
	f.Println(runtime.NumCPU() + 1)
	// Prints the number of CPUs I have + 1
}