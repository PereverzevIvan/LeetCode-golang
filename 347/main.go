package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}

	for _, v := range nums {
		hash[v]++
	}

	var numSet []int
	for k := range hash {
		numSet = append(numSet, k)
	}

	sort.SliceStable(numSet, func(i, j int) bool { return hash[numSet[i]] > hash[numSet[j]] })
	return numSet[:k]
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
