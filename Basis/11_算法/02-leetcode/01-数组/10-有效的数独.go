package main

import "fmt"

/*
有效的数独
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2f9gg/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	sudoku := isValidSudoku(board)
	fmt.Println(sudoku)
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		m1 := make(map[byte]bool)
		m2 := make(map[byte]bool)
		m3 := make(map[byte]bool)
		fmt.Printf("i: %d, num[i]: %v\n", i, board[i])

		// 判断每一行是否重复
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				fmt.Printf("j1: %d, num[j]: %v\n", j, board[i][j])
				if m1[board[i][j]] {
					return false
				}
				m1[board[i][j]] = true
			}

			// 判断每一列是否重复
			if board[j][i] != '.' {
				fmt.Printf("j2: %d, num[j]: %v\n", j, board[j][i])
				if m2[board[j][i]] {
					return false
				}
				m2[board[j][i]] = true
			}

			// 判断9宫格内的数据是否重复
			row := (i%3)*3 + j%3
			col := (i/3)*3 + j/3
			if board[row][col] != '.' {
				fmt.Printf("j3: %d, num[j]: %v\n", j, board[row][col])
				if m3[board[row][col]] {
					return false
				}
				m3[board[row][col]] = true
			}
		}
	}
	return true
}
