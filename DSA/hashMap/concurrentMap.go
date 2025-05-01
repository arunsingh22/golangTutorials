package main

import "sync"

// NOTE: Highly optimized for concurrent ready heavy access and low write.
// The sync.Map key is any type but it MUST always be comparable in nature.
var cmap = sync.Map{}

func Add(key string, value int) {
	cmap.Store(key, value)
}

func Get(key string) (int, bool) {
	if value, ok := cmap.Load(key); ok {
		return value.(int), true
	}
	return 0, false
}

func Delete(key string) {
	cmap.Delete(key)
}

func Clear() {
	cmap.Range(func(key, value interface{}) bool {
		cmap.Delete(key)
		return true
	})
}
func one() {
	Add("key1", 1)
	Add("key2", 2)
	Add("key3", 3)

	if value, ok := Get("key2"); ok {
		println("key2:", value)
	} else {
		println("key2 not found")
	}
	Delete("key2")
	if value, ok := Get("key2"); ok {
		println("key2:", value)
	} else {
		println("key2 not found")
	}
	Clear()
	if value, ok := Get("key1"); ok {
		println("key1:", value)
	} else {
		println("key1 not found")
	}
	if value, ok := Get("key3"); ok {
		println("key3:", value)
	} else {
		println("key3 not found")
	}
	// loadAndDelete
	// loadAndStore
	//
}
