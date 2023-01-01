package main

import "fmt"

// This interface helps in polymorphing the 3 types of caches
type evictionAlgorithm interface {
	evict()
}

// concrete type aka strategyClass 1
type lru struct{}

func (l *lru) evict() {
	fmt.Println("Eviction usign LRU caching")
}

// concrete type aka strategyClass 2
type lfu struct{}

func (l *lfu) evict() {
	fmt.Println("Eviction usign LFU caching")
}

// concrete type aka strategyClass 3
type fifo struct{}

func (l *fifo) evict() {
	fmt.Println("Eviction usign fifo caching")
}

// Class using the cache interface.
type cache struct {
	name      string
	evictAlgo evictionAlgorithm
}

// Public constructor
func NewCache(name string, algo evictionAlgorithm) *cache {
	c := &cache{
		name:      name,
		evictAlgo: algo,
	}
	return c
}

func (c *cache) setEvictionAlgorithm(e evictionAlgorithm) {
	c.evictAlgo = e
}

func (c *cache) display() {
	fmt.Println(c.name, c.evictAlgo)
}

// This is the client code which can use any caching algorithm without changing itself.
func main() {
	cache := NewCache("mycache", &lru{})
	cache.display()
	cache.evictAlgo.evict()
	cache.setEvictionAlgorithm(&fifo{})
	cache.evictAlgo.evict()

}
