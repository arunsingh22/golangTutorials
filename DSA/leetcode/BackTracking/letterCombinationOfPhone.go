package main

func letterCombinations(digits string) []string {
	result := []string{}

	if len(digits) == 0 {
		return result
	}
	dict := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	findMapping(dict, 0, digits, "", &result)
	return result
}

func findMapping(dict map[string]string, idx int, digits, curr string, result *[]string) {
	if idx == len(digits) {
		*result = append(*result, curr)
		return
	}
	charSet := dict[string(digits[idx])]
	for k := 0; k < len(charSet); k++ {
		curr += string(charSet[k])
		findMapping(dict, idx+1, digits, curr, result)
		curr = curr[:len(curr)-1]
	}
}
