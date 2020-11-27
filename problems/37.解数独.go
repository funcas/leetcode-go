package problems

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] 解数独
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for c := '1'; c <= '9'; c++ {
					if isValid(board, i, j, byte(c)) {
						board[i][j] = byte(c)
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row int, col int, c byte) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] != '.' && board[i][col] == c {
			return false
		}
		if board[row][i] != '.' && board[row][i] == c {
			return false
		}
		if board[row/3*3+i/3][col/3*3+i%3] != '.' &&
			board[row/3*3+i/3][col/3*3+i%3] == c {
			return false
		}
	}

	return true
}

// @lc code=end
