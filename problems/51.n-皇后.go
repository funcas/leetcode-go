package problems

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	ans := [][]int{}
	cols := make(map[int]struct{})
	pie := make(map[int]struct{})
	na := make(map[int]struct{})
	if n < 1 {
		return [][]string{}
	}
	var backtrace func(row int, path []int)
	backtrace = func(row int, path []int) {
		if row >= n {
			res := make([]int, n)
			copy(res, path)
			ans = append(ans, res)
			return
		}
		for col := 0; col < n; col++ {
			if _, ok := cols[col]; ok {
				continue
			}
			if _, ok := pie[col+row]; ok {
				continue
			}
			if _, ok := na[row-col]; ok {
				continue
			}
			cols[col] = struct{}{}
			pie[row+col] = struct{}{}
			na[row-col] = struct{}{}
			backtrace(row+1, append(path, col))
			// path = path[:len(path)-1]
			delete(cols, col)
			delete(pie, row+col)
			delete(na, row-col)
		}
	}
	backtrace(0, []int{})
	return formatAns(n, ans)
}

func formatAns(n int, ans [][]int) [][]string {
	board := [][]string{}
	for i := 0; i < len(ans); i++ {
		tmp := []string{}
		for j := 0; j < len(ans[i]); j++ {
			a := ""
			for k := 0; k < n; k++ {

				if k == ans[i][j] {
					a += "Q"
				} else {
					a += "."
				}
			}
			tmp = append(tmp, a)
		}
		board = append(board, tmp)
	}
	return board
}

// @lc code=end
