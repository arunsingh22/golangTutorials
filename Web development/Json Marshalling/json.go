package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"course"`
	Price    int      `json:"price"`
	Tags     []string `json:"Tags,omitempty"`
	Password string   `json:"-"`
}

func jsonEncode() {
	data := []course{
		{
			Name:  "GoLang course",
			Price: 124,
			Tags: []string{
				"dev",
				"test",
			},
			Password: "45#Hurgeon",
		},
		{
			Name:  "Nodejs course",
			Price: 1224,
			Tags: []string{
				"dev",
				"DevOps",
			},
			Password: "45#",
		},
	}

	//JSON Marshalling
	jsonOutput, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonOutput))
}

func jsonDecode() {
	var mycourse course
	//dummy JSON data
	dummyJSON := []byte(`
	{
    		"course": "GoLang course",
    		"price": 124,
    		"Tags": ["dev","test"]
  		}
	`)
	//first check if json is in right format or not.
	if json.Valid(dummyJSON) == true {
		json.Unmarshal(dummyJSON, &mycourse)
		fmt.Printf("%#v", mycourse) // Prints in right format
		// fmt.Println(mycourse) //does print what i want.
	} else {
		fmt.Println("Invalid JSON format")
	}

}

func main() {
	// jsonEncode()
	jsonDecode()
}
