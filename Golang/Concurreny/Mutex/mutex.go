package main

import "sync"

type balance struct {
	amt  int
	lock sync.RWMutex
}

func main() {

	lock.
	bal := &balance{
		amt : 10,
	}
	go func add(bal *struct,amt int ){
		lock.Lock()
		
	}(bal,5)
}
