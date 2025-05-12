package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Singleton is a creational design pattern, which ensures that only one object of its
// kind exists and provides a single point of access to it for any other code.

// Singleton has almost the same pros and cons as global variables.
// Although they’re super-handy, they break the modularity of your code.

type student struct{}

var singleInstance *student

var mu = &sync.Mutex{}

func getInstance() *student {
	if singleInstance == nil {
		mu.Lock()
		defer mu.Unlock()
		// check again
		if singleInstance == nil {
			fmt.Println("New Instance created..")
			singleInstance = &student{}
		}
		fmt.Println("Instance already created..")
	}
	fmt.Println("Instance already created..")
	return singleInstance
}

func main() {
	// for range 30 {
	// 	go getInstance()
	// }

	for range 10 {
		go getDBInstance(context.Background(), "mongo")
	}
	time.Sleep(2 * time.Second)
}
