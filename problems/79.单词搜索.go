package problems

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func Exist(board [][]byte, word string) bool {
	m, n, z := len(board), len(board[0]), len(word)
	var dfs func(i, j, k int, visited [][]int) bool

	dfs = func(i, j, k int, visited [][]int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == z-1 {
			return true
		}
		visited[i][j] = 1
		dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
		result := false
		for _, d := range dir {
			ii, jj := i+d[0], j+d[1]
			if ii >= 0 && ii < m && jj >= 0 && jj < n {
				if visited[ii][jj] == 0 {
					if dfs(ii, jj, k+1, visited) {
						result = true
						break
					}
				}
			}
		}
		visited[i][j] = 0
		return result
	}

	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
