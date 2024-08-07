package main

import (
	"fmt"
	"unicode"
)

func clearString(s string) string {
	var newS []rune

	for _, v := range s {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			newS = append(newS, unicode.ToLower(v))
		}
	}

	return string(newS)
}

func isPalindrome(s string) bool {
	newS := clearString(s)
	l, r := 0, len(newS)-1

	for l < r {
		if newS[l] != newS[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("race a car"))
}
