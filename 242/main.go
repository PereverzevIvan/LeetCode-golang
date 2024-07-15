package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := map[rune]int{}
	mapT := map[rune]int{}

	for _, c := range s {
		mapS[c] += 1
	}

	for _, c := range t {
		mapT[c] += 1
	}

	for _, sChar := range s {
		if mapS[sChar] != mapT[sChar] {
			return false
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	ss := strings.Split(s, "")
	st := strings.Split(t, "")

	sort.Strings(ss)
	sort.Strings(st)

	return strings.Join(ss, "") == strings.Join(st, "")
}

func main() {
	fmt.Println(isAnagram("cat", "rat"))
}
