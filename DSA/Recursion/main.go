package main

import "fmt"

func main() {
	// printingNtoOne(5)
	// printingOneToN(5)

	// findXInArr(1, 0, []int{1, 2, 3, 5, 7, 0, -1})
	// fmt.Println(findMinInArr(0, []int{1, 2, 3, 4, 5, 6, -7}))

	// Check if the array isSorted or not.
	// fmt.Println(checkIfArrIsSorted(0, []int{1, 2, 3, 4, 5, -6}))

	// factorial of a number
	// fmt.Println(fact(6))

	// fmt.Println(fibonnaci(6))

	// arr := []int{1, 2, 3, 4, 5}
	// // reverseArray(0, len(arr)-1, arr)
	// reverseArray2(0, arr)
	// fmt.Println(arr)

	str := "arun"
	fmt.Println(isPalindrome(0, len(str)-1, str))
}

// head recursion: post evaluation
func printingOneToN(n int) {
	// BC
	if n == 1 {
		fmt.Println(n)
		return
	}
	printingOneToN(n - 1)
	fmt.Println(n)
}

// tail recursion
func printingNtoOne(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printingNtoOne(n - 1)
}

func findXInArr(x, idx int, arr []int) {
	if len(arr)-1 < idx {
		fmt.Println("Not found!")
		return
	}
	if arr[idx] == x {
		fmt.Println("Found at index:", idx)
		return
	}
	findXInArr(x, idx+1, arr)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	n *= fact(n - 1)
	return n
}

func fibonnaci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	n = fibonnaci(n-1) + fibonnaci(n-2)
	return n
}

// TC: O(log(min(a, b)))
// SC == TC due to stack
func gcdRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdRecursive(b, a%b)
}

// SC is O(1)
func gcdIterative(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func checkIfArrIsSorted(i int, arr []int) bool {
	if i >= len(arr)-1 {
		return true
	}
	if arr[i] > arr[i+1] {
		return false
	}
	return checkIfArrIsSorted(i+1, arr)
}

// post-order evalution
func findMinInArr(idx int, arr []int) int {
	// Base condition
	if idx == len(arr)-1 {
		return arr[idx]
	}
	restMin := findMinInArr(idx+1, arr)
	if restMin < arr[idx] {
		return restMin
	}
	return arr[idx]
}

// func sumOfDigits(n int) int{
// 	// Base condition
// 	if n ==0 {
// 		return 0
// 	}
// 	sumOfDigits(n/10)
// 	unitNum := n%10
// }
