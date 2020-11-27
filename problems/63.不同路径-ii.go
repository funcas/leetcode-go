package problems

/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if obstacleGrid[0][j] != 1 {
					dp[i][j] = 1
				}
				continue
			}
			if j == 0 {
				if obstacleGrid[i][j] != 1 {
					dp[i][j] = 1
				}
				continue
			}
			if i > 0 && j > 0 {
				if obstacleGrid[i][j] == 1 {
					continue
				}
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, m)
	if obstacleGrid[0][0] != 1 {
		dp[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[j][i] == 1 {
				dp[j] = 0
				continue
			}
			if j >= 1 && obstacleGrid[j][i] == 0 {
				dp[j] = dp[j-1] + dp[j]

			}
		}
	}
	return dp[m-1]
}
