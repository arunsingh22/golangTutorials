package main

func reverseArray(i, j int, arr []int) {
	if i >= j {
		return
	}
	// swap the values
	arr[i], arr[j] = arr[j], arr[i]
	reverseArray(i+1, j-1, arr)
}

func reverseArray2(i int, arr []int) {
	if i > len(arr)/2 {
		return
	}
	// swap the values
	arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	reverseArray2(i+1, arr)
}
