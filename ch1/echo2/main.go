// Echo2 prints its command-line arguments.
// $ go run main.go Anything You Would Like main.go to Print Out

package main

import (
	"fmt"
	// The os package provides functions and values
	// that deal with the operating system.
	"os"
)

func main() {
	s, sep := "", ""
	// If we want to loop from the start / to the end,
	// we can use os.Args[:n] / os.Args[n:]
	for i, arg := range os.Args[1:] {
		// os.Args accesses command-line arguments
		// which returns a slice of strings.

		// os.Args[0] is the name of the command itself.
		// In this case, we want to print anything after the command.
		// so we want os.Arg[1:len(os.Args)] with spaces in between.
		s += sep + arg
		sep = " "

		// Exercise 1.2: Modify the echo program to print the index and value
		// of each of its arguments, one per line.
		fmt.Println(i+1, arg)
	}
	fmt.Println(s)

	// Exercise 1.1: Modify the echo program to also print os.Args[0]
	fmt.Println(os.Args[0] + " " + s)
}
