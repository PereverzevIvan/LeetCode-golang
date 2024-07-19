package main

import "fmt"

func hasDuplicates(arr []byte) bool {
	hash := make(map[byte]bool)

	for _, v := range arr {
		if _, inMap := hash[v]; inMap && v != byte(46) {
			return true
		} else {
			hash[v] = true
		}
	}

	return false
}

func isValidSudoku(board [][]byte) bool {
	// Проверка строк на наличие повторяющихся элементов
	for _, v := range board {
		if hasDuplicates(v) {
			return false
		}
	}

	// Проверка столбцов на наличие повторяющихся элементов
	for i := 0; i < len(board); i++ {
		var col []byte
		for j := 0; j < len(board[i]); j++ {
			col = append(col, board[j][i])
		}
		if hasDuplicates(col) {
			return false
		}
	}

	// Проверка блоков 3х3 на наличие повторяющихся элементов
	for row := 0; row < len(board); row += 3 {
		for col := 0; col < len(board[row]); col += 3 {
			var arr []byte
			arr = append(arr, board[row][col:col+3]...)
			arr = append(arr, board[row+1][col:col+3]...)
			arr = append(arr, board[row+2][col:col+3]...)
			fmt.Println(arr)

			if hasDuplicates(arr) {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(byte(' '))
}
