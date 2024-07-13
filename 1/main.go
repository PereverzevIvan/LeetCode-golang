package main

import "fmt"

/*
Дан массив чисел nums и целое число target. Нужно вернуть пару индексов элементов, сумма которых равна target
Можно считать, что каждый кейс имеет только одно решение.
*/
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	hash := map[int]int{} // Хэш таблица, в которой хранятся числа и индесы

	for i, num := range nums {
		if j, inMap := hash[target-num]; inMap {
			result = []int{i, j}
			break
		} else {
			hash[num] = i
		}
	}

	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

