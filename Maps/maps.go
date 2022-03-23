// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	//composite literal
	m := map[string]int{
		"arun":  23,
		"singh": 35,
	}
	//fmt.Printf("%v%T", m, m)

	//,ok idiom to check element is present of not
	if _, ok := m["arun"]; ok {
		fmt.Println("Hi is present")
	} else {
		fmt.Println("Absent")
	}

	//Deletes the  element in map
	delete(m, "arun")

	if _, ok := m["arun"]; ok {
		fmt.Println("Hi is present")
	} else {
		fmt.Println("Absent")
	}
}
