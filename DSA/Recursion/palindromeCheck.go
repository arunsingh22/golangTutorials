package main

func isPalindrome(i, j int, s string) bool {
	if i >= j {
		return true
	}
	if s[i] != s[j] {
		return false
	}
	return isPalindrome(i+1, j-1, s)
}
