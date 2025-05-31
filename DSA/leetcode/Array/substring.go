package main

import "fmt"

func generateSubstring(s string) {
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(s[i : j+1])
		}
	}
}
