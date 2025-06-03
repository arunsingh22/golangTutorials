package main

import (
	"slices"
	"strings"
)

// https://leetcode.com/problems/sort-characters-by-frequency/

type pair struct {
	ch  rune
	cnt int
}

func frequencySort(s string) string {
	dict := map[rune]int{}
	for _, ch := range s {
		dict[ch]++
	}

	ss := make([]pair, 0, len(dict))
	for k, v := range dict {
		ss = append(ss, pair{
			ch:  k,
			cnt: v,
		})
	}

	// custom sorting
	slices.SortFunc(ss, func(i, j pair) int {
		if i.cnt > j.cnt {
			return -1
		} else if i.cnt < j.cnt {
			return 1 // i has lower count, so i comes after j
		}
		return 0 // counts are equal
	})

	var ans strings.Builder
	for _, val := range ss {
		ans.WriteString(strings.Repeat(string(val.ch), val.cnt))
	}
	return ans.String()
}
