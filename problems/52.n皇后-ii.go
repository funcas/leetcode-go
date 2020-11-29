package problems

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N皇后 II
 */

// @lc code=start
func totalNQueens(n int) int {
	if n <= 0 {
		return 1
	}
	var cols []int
	var ldiagonal []int
	var rdiagonal []int
	var ans [][]int
	var backtrace func(row int, ret []int)
	backtrace = func(row int, ret []int) {
		if row >= n {
			ans = append(ans, ret)
			return
		}

		for col := 0; col < n; col++ {
			if contain(cols, col) || contain(ldiagonal, col+row) || contain(rdiagonal, col-row) {
				continue
			}
			cols = append(cols, col)
			ldiagonal = append(ldiagonal, col+row)
			rdiagonal = append(rdiagonal, col-row)

			backtrace(row+1, append(ret, col))

			cols = cols[:len(cols)-1]
			ldiagonal = ldiagonal[:len(ldiagonal)-1]
			rdiagonal = rdiagonal[:len(rdiagonal)-1]
		}
	}
	backtrace(0, []int{})
	return len(ans)
}

func contain(sl []int, item int) bool {
	for _, v := range sl {
		if item == v {
			return true
		}
	}
	return false
}

// 位运算解法，思路来自谭超《算法通关40讲》
func totalNQueens2(n int) int {
	if n <= 0 {
		return 1
	}
	var dfs func(row int, cols int, p int, q int)
	count := 0
	dfs = func(row int, cols int, p int, q int) {
		if row >= n {
			count++
			return
		}
		bits := ^(cols | p | q) & ((1 << n) - 1)
		for bits > 0 {
			tmp := bits & -bits
			dfs(row+1, tmp|cols, (p|tmp)<<1, (q|tmp)>>1)
			bits &= bits - 1
		}
	}
	dfs(0, 0, 0, 0)
	return count
}

// @lc code=end
