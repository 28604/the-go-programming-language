package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))

	// Exercise 1.1: Modify the echo program to also print os.Args[0]
	fmt.Println(os.Args[0] + " " + strings.Join(os.Args[1:], " "))

	// Exercise 1.2: Modify the echo program to print the index and value
	// of each of its arguments, one per line.
	for i, arg := range os.Args[1:] {
		fmt.Println(i+1, arg)
	}
}
