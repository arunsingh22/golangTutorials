package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func even(wg *sync.WaitGroup, syncChannel chan bool, done int32) {
	// prints even numbers.
	defer wg.Done()

	for i := 0; i < 10; i += 2 {
		fmt.Println("From even:", <-syncChannel)
		// // time.Sleep(time.Second)
		fmt.Printf("A: %d\n", i)
		syncChannel <- true
	}
	atomic.StoreInt32(&done, 1)
}

func odd(wg *sync.WaitGroup, syncChannel chan bool, done int32) {
	// prints odd numbers.
	defer wg.Done()

	for i := 1; i < 10; i += 2 {
		fmt.Println("From odd:", <-syncChannel)
		// time.Sleep(time.Second)
		fmt.Printf("B: %d\n", i)

		if atomic.LoadInt32(&done) != 0 {
			return
		}

		syncChannel <- true
	}
}

func main() {
	var done int32
	syncChannel := make(chan bool) // unbuffered channel.

	// go func() {
	// 	syncChannel <- true
	// }()
	// fmt.Println(<-syncChannel)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go even(wg, syncChannel, done)
	syncChannel <- true
	go odd(wg, syncChannel, done)

	wg.Wait()
	close(syncChannel)
}
