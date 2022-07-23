package main

import (
	"fmt"
	"strings" //all string related methods
)
// In Go, a string is in effect a read-only slice of bytes. 
func main() {
	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// fmt.Println(sample)

	name := "Hello world!, I am arun singh####"
	ss := strings.Split(name,",")
	fmt.Println(ss)

	sx := strings.Trim(name,"#")
	fmt.Println(sx)

	fmt.Println(string(name[4])) //this by default gives me UTF-8 representation of char o
}



package main

import (
	"fmt"
	"math"
	"strings"
)

// type _string struct {
// 	elements *byte // underlying bytes (immutable and unaddressable)
// 	len      int   // number of bytes
// }

type base struct {
	a *string
	b int
}

func reverseString(s []byte) {
	//Declaring more than 2 variables in for-loop
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		//basic swap operations
		s[i], s[j] = s[j], s[i]
	}
}

func mostWordsFound(sentences []string) int {
	var ans, cnt float64
	for _, str := range sentences {
		cnt = 0
		// strings.Fields(str) splits the strings on whitespace and \t and
		// returns a slice
		for range strings.Split(str, " ") {
			cnt++
		}
		ans = math.Max(cnt, ans)
	}

	return int(ans)
}

func display(str string) {
	for _, k := range str {
		fmt.Println(string(k))
	}
}

func main() {
	// arr := []string{"arun", "singh", "zsh"}

	// //Sorts the slice
	// sort.Slice(arr, func(i, j int) bool {
	// 	return arr[i] < arr[j]
	// })

	// sort.SliceStable(arr, func(i, j int) bool {
	// 	return arr[i] > arr[j]
	// })

	// for _, k := range arr {
	// 	fmt.Println(k)
	// }

	// fmt.Println(&s[0]) //taking address of any char or byte is NOT allowed.
	display("arun singh")

	dict := make(map[string]string)
	fmt.Printf("Dict :%s and length of dict: %d", dict, len(dict))

}
