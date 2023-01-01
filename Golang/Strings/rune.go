// Go program to illustrate how to check the
// given rune is a Unicode punctuation
// character or not
package main

import (
	"fmt"
	"unicode"
)

// Main function
func main() {

	// Creating a slice of rune
	val := []rune{'g', 'f', 'G', '#', ',', ':'}

	// Checking each element of the given slice
	// of the rune is a Unicode punctuation
	// character or not
	// Using IsPunct() function
	for i := 0; i < len(val); i++ {

		if unicode.IsPunct(val[i]) == true {

			fmt.Println("It is a Unicode punctuation character")

		} else {

			fmt.Println("It is not a Unicode punctuation character")
		}
	}
}
