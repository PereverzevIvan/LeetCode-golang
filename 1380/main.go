package main

import "slices"

func isMin(arr []int, num int) bool {
	minimalNum := slices.Min(arr)
	return minimalNum == num
}

func isMax(arr []int, num int) bool {
	maxNum := slices.Max(arr)
	return maxNum == num
}

func getColValues(matrix [][]int, col int) []int {
	var result []int
	for i := 0; i < len(matrix); i++ {
		result = append(result, matrix[i][col])
	}

	return result
}

func luckyNumbers(matrix [][]int) []int {
	var result []int

	for _, row := range matrix {
		for j, value := range row {
			if isMin(row, value) {
				colValues := getColValues(matrix, j)
				if isMax(colValues, value) {
					result = append(result, value)
				}
			}
		}
	}

	return result
}
