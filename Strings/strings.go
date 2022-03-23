package main

import (
	"fmt"
	"strings" //all string related methods
)
// In Go, a string is in effect a read-only slice of bytes. 
func main() {
	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// fmt.Println(sample)

	name := "Hello world!, I am arun singh####"
	ss := strings.Split(name,",")
	fmt.Println(ss)

	sx := strings.Trim(name,"#")
	fmt.Println(sx)

	fmt.Println(string(name[4])) //this by default gives me UTF-8 representation of char o
}
