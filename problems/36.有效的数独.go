package problems

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 9)
	cols := make([]map[byte]struct{}, 9)
	boxs := make([]map[byte]struct{}, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]struct{})
		cols[i] = make(map[byte]struct{})
		boxs[i] = make(map[byte]struct{})
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			item := board[row][col]
			if _, ok := rows[row][item]; ok {
				return false
			}
			if _, ok := cols[col][item]; ok {
				return false
			}
			if _, ok := boxs[col/3*3+row/3][item]; ok {
				return false
			}

			rows[row][item] = struct{}{}
			cols[col][item] = struct{}{}
			boxs[col/3*3+row/3][item] = struct{}{}
		}
	}
	return true
}

// @lc code=end

func isValidSudoku2(board [][]byte) bool {
	rows := [9]int{}
	cols := [9]int{}
	boxes := [9]int{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			item := board[i][j]
			boxIdx := i/3*3 + j/3

			if (rows[i]>>item&1) == 1 ||
				(cols[j]>>item&1) == 1 || (boxes[boxIdx]>>item&1) == 1 {
				return false
			}

			rows[i] = rows[i] | (1 << item)
			cols[j] = cols[j] | (1 << item)
			boxes[boxIdx] = boxes[boxIdx] | (1 << item)
		}
	}
	return true
}
