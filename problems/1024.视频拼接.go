package problems

/*
 * @lc app=leetcode.cn id=1024 lang=golang
 *
 * [1024] 视频拼接
 */

// @lc code=start
func videoStitching(clips [][]int, T int) int {
	dp := make([]int, T+1)
	for i := range dp {
		dp[i] = 1<<31 - 1
	}
	dp[0] = 0
	for i := 1; i <= T; i++ {
		for j := 0; j < len(clips); j++ {
			l, r := clips[j][0], clips[j][1]
			if l < i && i <= r && dp[l]+1 < dp[i] {
				dp[i] = dp[l] + 1
			}
		}
	}
	if dp[T] == 1<<31-1 {
		return -1
	}
	return dp[T]
}

// @lc code=end
