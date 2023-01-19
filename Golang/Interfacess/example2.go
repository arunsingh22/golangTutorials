package main

import "fmt"

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {
	var x interface{} = DummyWriter{}

	// I can only call Write after extracting out the concrete type of x.
	// This idea makes more sense after reading example 3.
	a, err := x.(DummyWriter).Write([]byte("12"))
	if err != nil {

	}
	fmt.Println(a)

}
