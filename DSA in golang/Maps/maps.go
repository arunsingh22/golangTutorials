// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	dict := map[string]int{} // nil, no physical storage only declaration of type.
	fmt.Println("First: ", dict)
	// dict["x"] = 10          // ERROR

	//Note: Recommeded way because maps can be read even if they
	// are nil but if inserted then it will panic.
	dict2 := make(map[string]int, 10) //non-nil but empty because now there is a size 10 hashtable getting created under the hood.
	dict2["peacock"] = 10
	fmt.Println("length of dict2: ", len(dict2))

	// make(map[string]int) equilvalent to map[string]int{}
	dict3 := map[string]int{} // non-nil BUT empty
	dict3["key"] = 1
	fmt.Println(dict3)

	fmt.Println("--------RANGE OVER MAP----------")
	//composite literal
	m := map[string]int{
		"arun":   23,
		"singh":  35,
		"shilpa": 53,
		"yadav":  90,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("------------------")

	//Maps have a special two-result lookup function called , ok idiom
	if val, ok := m["arun"]; ok {
		fmt.Println("Present: ", val)
	} else {
		fmt.Println("Absent")
	}

	//Deletes the  element in map
	delete(m, "arun")

	//Note: Maps can't be compared to one another; maps can be compared only to nil as a special case.
	// dict == dict2 SYNTAX ERROR

}
