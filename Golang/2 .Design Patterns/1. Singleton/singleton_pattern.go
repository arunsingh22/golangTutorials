package main

import (
	"fmt"
	"sync"
)

// https://medium.easyread.co/just-call-your-code-only-once-256f69ed39a8
type database struct{}

var (
	doOnce sync.Once
	dbConn *database
	lock   sync.RWMutex
)

// check-lock-check pattern for singleton
// This approch still has a potential to break in production under extreme stress as the if condition is not atomic in nature
func getConnection() *database {
	if dbConn == nil { //check
		lock.Lock()    //lock 
		defer lock.Unlock()
		// If condition is not atomic and hence there can be certain edge cases where this will fail.
		if dbConn == nil { //check
			fmt.Println("Creating Connection Object to database")
			dbConn = &database{}
		}
	}
	return dbConn
}

// Best way to implement singleton pattern in GO.
func getConnectionOnlyOnce() *database {
	doOnce.Do(func() {
		fmt.Println("Creating Connection Object to database by using sync.Once ")
		dbConn = &database{}
	})
	return dbConn
}

func ExampleOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
	// Output:
	// Only once
}

func main() {
	// var once sync.Once
	// fx := func() {
	// 	fmt.Println("One time")
	// }
	// for i := 0; i < 10; i++ {
	// 	once.Do(fx)
	// 	fmt.Println("other")
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(getConnectionOnlyOnce())
	// }

	ExampleOnce()
}
