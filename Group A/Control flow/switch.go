package main

import "fmt"

func main() {
	a := 3

	//Note: Unlike other programming lang, in go switch automatically breaks at
	// each case and default is the last one to execute if no other case matches.

	switch a {
	case 1, 2, 3:
		fmt.Println("The value of a:", a)
		fmt.Println("-------------------")
	case 4:
		fmt.Println("On case 4")
	case 5, 6:
		fmt.Println("On case 5")
		fmt.Println("On case 6")
	default:
		fmt.Println("on default")

	}

	// Second kind of switch statement.
	switch x := getValue(); x {
	case "a":
		fmt.Println("on a")
	default:
		fmt.Println("Reached on default")
	}
}

func getValue() string {
	return "a"
}
