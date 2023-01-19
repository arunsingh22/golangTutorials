package main

import "fmt"

// https://zetcode.com/golang/struct/

// Just like C, Go structs also have structural padding and packing concept.
// Anything declared using composite literal way has key:value, style of writing be it maps or struct
//creating a person type
type person struct {
	first string
	last  string
	age   int
}

//Embedded structs; they are APPROXIMATION TO INHERITANCE.
type secretAgent struct {
	person //person struct is embbedd into secretAgent
	rtk    bool
}

func main() {
	//composite literal to create a struct
	// p1 := person{
	// 	first: "arun",
	// 	last : "singh",
	// 	age : 25,
	// }

	// fmt.Println(p1.age,p1.first)

	sa1 := secretAgent{
		rtk: true,
		person: person{
			first: "ash",
			last:  "singh",
			age:   23,
		},
	}

	fmt.Println(sa1)
	fmt.Println(sa1.age, sa1.first, sa1.last, sa1.rtk) //type promotion has happend here

}
