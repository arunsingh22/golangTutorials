package main

import (
	"fmt"
	"sync"
	"time"
)

var memCount int

func main() {

	pool := &sync.Pool{
		New: func() interface{} {
			memCount++
			mem := make([]int, 1024)
			return &mem
		},
	}

	numWorkers := 1024 * 1024 
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			mem := pool.Get().(*[]int)
			time.Sleep(time.Second * 2)
			pool.Put(mem)
		}()
	}

	wg.Wait()
	fmt.Println("Number of workerPool:", memCount)
	fmt.Println("End Main")

}
