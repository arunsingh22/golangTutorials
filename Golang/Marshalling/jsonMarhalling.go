package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string `json:"-"` //explicity we don't want to jsonify
	Second      string `json:"family_name"`
	Age         int    `json: "age"`
	notexported int    //smallcase fields are not exported as json
}

// type Person struct {
// 	First  string
// 	Second string
// 	Age    int
// }

func (p Person) printName() string {

	return p.First + " " + p.Second
}

func main() {
	p := Person{
		"arun", "singh", 25,0,
	}
	//Marshalling the data.
	bs, _ := json.Marshal(p)
	// fmt.Println(bs)
	fmt.Println("----------------------")

	res := Person{}
	e := json.Unmarshal(bs, &res)
	fmt.Println(res.First, res.Age)
	if e != nil {

	}

}
