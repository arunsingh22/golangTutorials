# When are values constructre on the Heap in Go
---------------------------------------------------
1. When a value could **possibly** be referenced after the func that constructed the values returns.
2. When the compiler determines a values is too large to fit on the stack. 
3. When the compiler doesn't know the size of DS during compile time  eg : slice, map, Channels, Strings.-> these are not guranteed to get placed on heap.

- Escape Analysis in compile time behaviour
- Ask the compiler to find out , Don't guess 
    - **go build -gcflags "-m=2"**

------------------------------------------------------------------------
**CASE 1:**
func main() {
	// since the size is know and is statically declared slice will go on STACK frame.
	b := make([]byte,10)
	read(b) 
	fmt.Print(b)
}
func read(b []byte) {
	//writing into the slice
}

-------------------------------------------------------------------------
**CASE 2:**
func main() {
	b := read() 
	// b holds a reference to []byte ie, []byte will be allocated on heap.
	fmt.Print(b)
}
func read() []byte {
	b := make([]byte, 10)
	return b
}
=========================================================
**The above example EXPLAINS io.Reader**
type Reader interface{
    Read(b []byte) (n int, err error)
}

Instead of: This will create alot of load for GC
type Reader interface{
    Read(n int) (b []byte, err error)
}
========================================================

