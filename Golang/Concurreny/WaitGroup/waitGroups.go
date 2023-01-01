package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//step 1: Instanitate the wg
	var wg sync.WaitGroup
	// wg := new(sync.WaitGroup)

	now := time.Now()
	// wg.Add(10) // step 2: Add all the goroutines for which you want to wait into the waitGroup struct
	for i := 0; i < 10; i++ {
		wg.Add(1) // individual adding is littel slowwwww only use when u don't know in advance #workers
		go worker(&wg, i)
	}
	fmt.Println("Number of Goroutines: ", runtime.NumGoroutine())
	fmt.Println(time.Since(now))
	wg.Wait() // step 4: wg.Wait would wait for all the goroutines.
	fmt.Println("Waiting for all the goroutines...Done!")
}

// NOTE: wg are ALWAYS PASSED by REFERENCE, if not then deadlock happens
func worker(wg *sync.WaitGroup, id int) {
	defer wg.Done() // step 3: wg.Done would reduce the count of goroutines value by 1
	time.Sleep(time.Millisecond * 10)
	fmt.Println("Worker: ", id)
}
