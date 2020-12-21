package problems

/*
 * @lc app=leetcode.cn id=1137 lang=golang
 *
 * [1137] 第 N 个泰波那契数
 */

// @lc code=start
func Tribonacci(n int) int {
	dp0, dp1, dp2 := 0, 1, 1
	ans := 0
	if n == 1 {
		ans = dp1
	}
	if n == 2 {
		ans = dp2
	}
	for i := 3; i <= n; i++ {
		ans = dp0 + dp1 + dp2
		dp0, dp1, dp2 = dp1, dp2, ans
	}
	return ans
}

// @lc code=end
