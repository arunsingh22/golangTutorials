package main

// | Type                           | Needs pointer to update caller? | Why?                                             |
// | ------------------------------ | ------------------------------- | ------------------------------------------------ |
// | `map`                          | ❌ No                            | Internally shared, all changes are reflected     |
// | `slice` (content change)       | ❌ No                            | Underlying array is shared                       |
// | `slice` (append/change length) | ✅ Yes                           | Append may reallocate, losing original reference |

// Map is always passed as ref so the changes will ALWAYS reflect back in the caller
// In slices things are pretty ticky

func main() {
	// slice := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Printf("slice len %d and cap %d\n", len(slice), cap(slice))
	// modifyTheUnderlyingSlice(slice)
	// fmt.Printf("slice len %d and cap %d\n", len(slice), cap(slice))
	// fmt.Println(slice)

	// addNewItemsToSlice(slice)
	// fmt.Println("After addNewItemsToSlice: ", slice)

	// printSlice(slice[2:4])
	GuessOutput2()
	// fullSlicing()

}

func modifyTheUnderlyingSlice(slice []int) {
	slice[0] = 11
	slice[1] = 12
}
