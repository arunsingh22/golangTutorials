package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	str := "arun is a good boy,124!!#XYZ"
	fmt.Println(strings.Contains(str, "arun")) // checks for substrings
	fmt.Println(strings.HasPrefix(str, "ar"))
	fmt.Println(strings.HasSuffix(str, "boy"))

	// doesn't modify the original string
	// Just like strings.ToUpper and strings.ToLower there is unicode pkg which has both
	// unicode.ToLower() & unicode.ToUpper() but only for runes and not for strings.
	fmt.Println(strings.ToUpper(str))

	fmt.Println(str)

	//-------------------------------------
	for _, c := range str {
		if unicode.IsUpper(c) == true {
			fmt.Print(string(c))
		}
		if unicode.IsLower(c) == true {
			fmt.Print(string(c))
		}
		if unicode.IsDigit(c) {
			fmt.Println(string(c))
		}
		if unicode.IsPunct(c) {
			fmt.Println(c)
		}

	}
	// fmt.Println()
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
