// Dup1 prints the text of each line that appears more than once
// in the standard input, preceded by its count.

// Programs for file copying, printing, searching, sorting, counting
// all have a similar structure: a loop over the input, some computation on each element,
// and generation of output on the fly or at the end.

package main

import (
	// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object.
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Make initializes an object of type slice, map, or chan.
	counts := make(map[string]int)
	// Scanner provides a convenient interface for reading data.
	// Input is a new Scanner that reads from the standard input.
	input := bufio.NewScanner(os.Stdin)
	// Scan advances the Scanner to the next token.
	// So, we keep scanning the input until it returns false upon hitting EOF or I/O error.
	// Each time scanning the token, we count the times a string occurred.
	// Text returns the string of the latest token (pass by value).
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential error from input.Err()

	for line, n := range counts {
		if n > 1 {
			// %d: decimal
			// %f, %g, %e: floating point
			// %t: boolean
			// %s: string
			// %q: quoted string
			// %v: any value in the natural format
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
