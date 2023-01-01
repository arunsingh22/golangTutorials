package benchmark

import (
	"sync"
	"testing"
)

func Add1(b *testing.B) {
	b.ReportAllocs()

	var count int
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()

}
