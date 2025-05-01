package main

import "testing"

var m = map[int]int{}

// Note: The fastest of all method
func Benchmark_increment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99]++
	}
}

func Benchmark_plusone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99] += 1
	}
}

func Benchmark_addition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m[99] = m[99] + 1
	}
}

// go test -bench=.
// go test -bench=. -count 5
// go test -bench=. -benchtime=10s
// go test -bench=. -benchtime=10s -benchmem
