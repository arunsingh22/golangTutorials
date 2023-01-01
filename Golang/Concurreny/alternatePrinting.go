//https://blog.logrocket.com/how-use-go-channels/

// output:
// A:0
// B:1
// A:2
// B:3
// A:4
// B:5

package main

import (
	"fmt"
	"sync"
	"time"
)

func A(ch1 chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 10; i = i + 2 {
		val := <-ch1
		if val == true {
			fmt.Println("A:", i)
			time.Sleep(time.Second * 1)
			go func() {
				ch1 <- false
			}()
		}
	}
	wg.Done()
}

func B(ch1 chan bool, wg *sync.WaitGroup) {
	for i := 1; i < 10; i = i + 2 {
		val := <-ch1 //false
		if val == false {
			time.Sleep(time.Second * 1)
			fmt.Println("B:", i)
			go func() {
				ch1 <- true
			}()
		}

	}
	wg.Done()
}

func main1() {

	ch1 := make(chan bool)
	// true : B can print
	// false : A can print
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		ch1 <- false
	}()

	go A(ch1, &wg)
	ch1 <- false
	go B(ch1, &wg)

	wg.Wait()
	close(ch1)
	fmt.Println("Done main")

}
