package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "luffy is still joyboy"

	// Split the string into words
	words := strings.Fields(str)
	// strings.Split(str, " ")

	// Loop over each word
	for _, word := range words {
		fmt.Println(word)
	}
}
