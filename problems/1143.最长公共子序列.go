package problems

/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			memo[i] = append(memo[i], -1)
		}
	}
	var lcs func(i int, j int) int
	lcs = func(i int, j int) int {
		if i == len(text1) || j == len(text2) {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		if text1[i] == text2[j] {
			memo[i][j] = 1 + lcs(i+1, j+1)
		} else {
			memo[i][j] = max(lcs(i+1, j), lcs(i, j+1))
		}
		return memo[i][j]
	}
	return lcs(0, 0)
}

// @lc code=end
