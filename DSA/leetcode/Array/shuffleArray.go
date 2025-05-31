package main

import (
	"math/rand"
)

func InbuiltShuffling(arr []int) {
	// This internally uses Fisher-yates algorithm
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

// Fisher-Yates shuffle
func shuffleArray(arr []string) {
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		randIdx := rand.Intn(i + 1)
		// swap the value
		arr[i], arr[randIdx] = arr[randIdx], arr[i]
	}
}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// arr := []string{"arun", "singh", "is", "a", "good", "boy"}
	// fmt.Println("Before Shuffling: ", arr)
	// shuffleArray(arr)
	// fmt.Println("After Shuffling: ", arr)

	s := "abc"
	generateSubstring(s)
}
