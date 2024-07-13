package main

import "fmt"

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)
	for _, num := range nums {
		if _, inMap := hash[num]; inMap {
			return true
		} else {
			hash[num] = true
		}
	}

	return false
}

func main() {
	a := []int{1, 3, 4, 2}
	fmt.Println(containsDuplicate(a))
}
