package main

import (
	"log"
	"time"
)

func main() {
	stopper := time.After(10 * time.Second) //ONE TIME TIMER
	ticker := time.NewTicker(1 * time.Second).C // REPEATED TIMER.
 
	for {
		select {
		case <-ticker:
			log.Println("tick")
		case <-stopper:
			log.Fatal("Done")
			// default:
			// 	log.Println("defualt") //using default makes infinite loop
		}

	}
}
