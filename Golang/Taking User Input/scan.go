package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	random_int := rand.Intn(2) // rand.Int() returns pseudo-random number from  seed to inf

	input := 0
	if _, err := fmt.Scanf("%d", &input); err != nil {
		fmt.Println("Wrong set of args:", err)
	}

	if input == random_int {
		fmt.Println("Match")
	} else {
		fmt.Println("random_int: ", random_int)
		fmt.Println("No Match")
	}

}
