package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('*')
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Welcome: ", name)
	}
}
