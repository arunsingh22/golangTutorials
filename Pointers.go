// https://stackoverflow.com/questions/5231096/time-complexity-of-power/8876837

package main

import "fmt"

func main() {

	a := 5
	b := &a
	fmt.Println(a, &a)

	if *b==a{
		fmt.Println("Equal:",*b)
	}

}
