package main

import (
	"fmt"
	"sort"
	"strings"
)

// https://leetcode.com/problems/group-anagrams/description/

// Функция для сортировки знаков внутри строки
func sortCharsInString(s string) string {
	ss := strings.Split(s, "")
	sort.Strings(ss)
	return strings.Join(ss, "")
}

// Функция для группировки анаграм в любом порядке
func groupAnagrams(strs []string) [][]string {
	var result [][]string
	hash := map[string][]string{}

	for _, s := range strs {
		ss := sortCharsInString(s)
		if _, inMap := hash[ss]; inMap {
			hash[ss] = append(hash[ss], s)
		} else {
			hash[ss] = []string{s}
		}
	}

	for _, v := range hash {
		result = append(result, v)
	}

	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
