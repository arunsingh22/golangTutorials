package main

// func main() {
// 	n := 1_000_000
// 	m := make(map[int]*[]byte)
// 	printAlloc()
// 	for i := 0; i < n; i++ { // Adds 1 million elements
// 		m[i] = &[]byte{'1'}
// 	}
// 	printAlloc()

// 	for i := 0; i < n; i++ { // Deletes 1 million elements
// 		delete(m, i)
// 	}
// 	runtime.GC() // Triggers a manual GC
// 	printAlloc()
// 	runtime.KeepAlive(m) // Keeps a reference to m so that the map isn’t collected

// }

// func printAlloc() {
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	fmt.Printf("%d MB\n", m.Alloc/(1024*1024))
// }
