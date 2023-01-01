package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0 // critical resource
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		count++
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		count = count - 2
	}(wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		count = count + 2
	}(wg)
	wg.Wait()

	fmt.Println("Count: ", count)
	//go run -race racing_due_waitgroup.go
	// ==================
	// WARNING: DATA RACE
	// Read at 0x00c0000140d8 by goroutine 8:
	//   main.main.func2()
	//       D:/GoLang/Group A/Concurreny/WaitGroup/racing_due_waitgroup.go:19 +0x7e
	//   main.main┬╖dwrap┬╖2()
	//       D:/GoLang/Group A/Concurreny/WaitGroup/racing_due_waitgroup.go:20
}
