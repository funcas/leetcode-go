package problems

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	dp_i_1, dp_i_2 := 1, 0
	dp_i := 1
	for i := 2; i <= N; i++ {
		dp_i = dp_i_1 + dp_i_2
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}

	return dp_i
}

// @lc code=end
