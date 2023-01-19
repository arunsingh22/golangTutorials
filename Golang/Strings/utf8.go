package main

import (
	"fmt"
)

func main() {
	x := "a😘"
	fmt.Println(x)

	hindi_str := "मेरा नाम अरुण सिंह है"
	fmt.Println(hindi_str)

	//NOTE: A simple for-loop iterates over BYTES and therefore indexing a string (using for loop on it) accesses individual bytes, not characters.
	for i := 0; i < len(hindi_str); i++ {
		fmt.Printf("%v -> %d -> %b\n", string(hindi_str[i]), hindi_str[i], hindi_str[i])
	}

	// for len(hindi_str) > 0 {
	// 	char, size := utf8.DecodeRune([]byte(hindi_str))
	// 	fmt.Printf("%c -> %3d -> %b size.\n", char, char, size)
	// 	hindi_str = hindi_str[size:]
	// }

	s := "Hellõ World👌"
	fmt.Println("Length: ", len(s)) // len calc ONLY the no of bytes and ie though the char count is 2 it prints 12 as len because õ takes 2 bytes and 👌takes 4 bytes so total = 16 bytes.

	// range within a for loop, range will return rune/int32 and byte index of the character.
	// A rune == codepoint == int32
	// UTF-8 converts codepoint <--> binary
	for index, char := range s {
		fmt.Printf("Char at index %d is %c -> %v:\n", index, char, int32(char))
	}

}
