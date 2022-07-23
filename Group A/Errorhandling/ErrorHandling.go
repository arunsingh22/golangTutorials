package ErrorHandling

import (
	"fmt"
	"log"
	"os"
)

func ExportedFunction() {
	fmt.Println("Inside Errorhandling pkg/Main")
	f, e := os.Create("log.txt")
	if e != nil {
		fmt.Println(e)
		log.Println(e)
	}
	defer f.Close()

	log.SetOutput(f)

	_, e = os.Open("no-file.txt")
	if e != nil {
		log.Println(e)
	}
}
