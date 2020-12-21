package problems

/*
 * @lc app=leetcode.cn id=861 lang=golang
 *
 * [861] 翻转矩阵后的得分
 */

// @lc code=start
func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])
	ans := m * (1 << (n - 1))
	for j := 1; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			if A[i][0] == 1 {
				ones += A[i][j]
			} else {
				ones += (1 - A[i][j]) // 翻转
			}
		}
		k := max(ones, m-ones)
		ans += k * (1 << (n - j - 1))
	}
	return ans

}

// @lc code=end
