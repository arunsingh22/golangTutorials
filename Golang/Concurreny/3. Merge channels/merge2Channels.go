package main

//https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308
import (
	"fmt"
	"time"
)

// Merging channels : https://medium.com/justforfunc/two-ways-of-merging-n-channels-in-go-43c0b57cd1de

func main() {
	a := asChan(1, 2, 3, 4, 5) // generator pattern
	b := asChan(6, 7, 8, 9)    //generator pattern
	c := merge(a, b)

	// Note: when ranging over channels only we get 1 value,no matter if the
	// channel is readonly or RW.
	for val := range c {
		// if !ok {
		// 	fmt.Println("channel closed")
		// }
		fmt.Println(val)
	}
}

// Method 1 : Works but consumes unnecessary CPU time as alot of waiting is done , we need a mechanism to make the channels INACTIVE after we are done using it.
// func merge(a, b <-chan int) <-chan int {
// 	outCh := make(chan int)
// 	go func() {
// 		aDone, bDone := false, false
// 		defer close(outCh) // close the channel as soon as done writing
// 		for aDone == false || bDone == false {
// 			select {
// 			case v, ok := <-a:
// 				if !ok {
// 					// fmt.Println("a is Done")
// 					aDone = true
// 					continue
// 				}
// 				outCh <- v
// 			case v, ok := <-b:
// 				if !ok {
// 					// fmt.Println("b is Done")
// 					bDone = true
// 					continue
// 				}
// 				outCh <- v
// 			}
// 		}
// 	}()
// 	return outCh
// }

// Method 2: We will employee nil channels to bascially DISABLE/INACTIVATE the channels in select block reading, writing or closing a nil channel blocks
// receiving from a nil channels blocks forever. So to disable a case receiving from a channel, we can simply set that channel to nil!
func merge(a, b <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					fmt.Println("a is Done")
					continue
				}
				outCh <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					fmt.Println("b is Done")
					continue
				}
				outCh <- v
			}
		}
	}()
	return outCh
}

func asChan(in ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch) // the writer always closes the channel
		for _, x := range in {
			ch <- x
			time.Sleep(time.Millisecond * 5)
		}

	}()
	return ch
}
