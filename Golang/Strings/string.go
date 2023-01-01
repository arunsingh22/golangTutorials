package main

import (
	"fmt"
	"math"
	"sort"
	"strings" //all string related methods
)

// type _string struct {
// 	elements *byte // underlying bytes (immutable and unaddressable)
// 	len      int   // number of bytes
// }

//1. In Go, a string is in effect a read-only slice of bytes.
//2. fmt.Println(&s[0]) //taking address of any char or byte is NOT allowed.

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
		// strings.Split(str) splits the strings on whitespace and \t and
		// returns a slice
		for range strings.Split(str, " ") {
			cnt++
		}
		ans = math.Max(cnt, ans)
	}

	return int(ans)
}

//ranging over strings
func display(str string) {
	for _, k := range str {
		fmt.Println(string(k)) //this by default gives me UTF-8 representation of char o
	}
}

func firstPalindrome(words []string) string {

	// anonymous helper func
	checkPalindrome := func(w string) bool {
		for i, j := 0, len(w)-1; i < j; i, j = i+1, j-1 {
			if w[i] != w[j] {
				return false
			}
		}
		return true
	}

	for _, word := range words {
		if checkPalindrome(word) == true {
			return word
		}
	}
	return ""
}

func main() {
	arr := []string{"arun", "singh", "zsh"}

	//Sorts the slice in increasing order
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	//Stable sorting in decreasing order
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	for _, k := range arr {
		fmt.Println(k)
	}

	dict := make(map[string]string)
	fmt.Printf("Dict :%s and length of dict: %d", dict, len(dict))

	//-----------------------------------------------
	fmt.Println("---- string inbuilt functions ------")
	fmt.Println("Count: ", strings.Count("arunsingh", "ar"))
	fmt.Println(strings.TrimSpace("   	arun singh    "))
	fmt.Println(strings.TrimLeft("   	arun singh  i  ", "\t \v"))

}
