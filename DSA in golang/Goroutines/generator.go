package generator


// Generator Pattern is useful for lazy evaluation—producing a sequence of values on 
// demand rather than computing them all upfront. It allows you to generate data on the fly, 
// especially for large datasets or infinite sequences, without holding everything in memory.


// producer 
func IntGenerator(limit int) <-chan int {
	ch := make(chan int)
	go func(){
		for i :=0;i<=limit ;i++{
			ch <- i
		}
	}()
	return ch
}

func main(){
	limit := 10

	// consumer de-coupled with producer 
	for i := range IntGenerator(limit){
		println(i)
	}
}