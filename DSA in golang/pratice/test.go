package main

import (
	"fmt"
	"sort"
	"time"
)

// arr = [1,5,3,1]
// o/p = 1 -> 0 & 3
//       3 -> 2
//       5 -> 1

// 1 - 0 
// 1 - 3 


//iterate
	// prev := 0
	// for i := 0; i < len(input); i++ {
	// 	sleepTime := input[i] - prev
	// 	time.Sleep(time.Duration(sleepTime) * time.Second)
	// 	prev = input[i]
	// 	fmt.Println(input[i], dict[input[i]])
	// 	// Check to keep i within limits
	// 	if i+1 < len(input) {
	// 		for input[i] == input[i+1] {
	// 			i++
	// 		}
	// 	}
	// }

func worker(duration int,idx int,chan int) chchan<- int {
	time.Sleep(time.Duration(duration) * time.Second) 
	fmt.Println(duration,idx)
	chan <-idx
}

func main(){
	input := []int{1, 5, 3, 3, 3}
	dict := map[int][]int{}

	for idx, num := range input {
		dict[num] = append(dict[num], idx)
	}

	//sorting
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	ch := make(chan int,len(input))
	for i,duration := range input{
		go worker(duration,i,ch)
	}

	fmt.Println(<- ch)

}
