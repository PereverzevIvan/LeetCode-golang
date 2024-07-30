package main

import (
	"fmt"
	"sort"
)

// Функция для получения множества из массива чисел
func getSet(nums []int) []int {
	var result []int
	hash := map[int]bool{}
	for _, v := range nums {
		hash[v] = true
	}

	for k := range hash {
		result = append(result, k)
	}
	return result
}

// Функция для получения длины самой длинной последовательности подряд идущих чисел
func longestConsecutive(nums []int) int {
	nums = getSet(nums)

	if len(nums) == 1 || len(nums) == 0 {
		return len(nums)
	}

	sort.Ints(nums)

	maxSequence := 0
	curSequence := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			curSequence++
		}

		if nums[i] != nums[i-1]+1 || i == len(nums)-1 {
			maxSequence = max(maxSequence, curSequence)
			curSequence = 1
		}
	}

	return maxSequence
}

func main() {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
