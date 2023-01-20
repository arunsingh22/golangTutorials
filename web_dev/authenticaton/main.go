package main 

import(
	"fmt"
	"encoding/base64"
)

func main{
	fmt.Println(base64.StdEncoding.Encode("username:password"))
	// curl -u username:pwd google.com -v -> Authorization:Basic 25AXayt
}