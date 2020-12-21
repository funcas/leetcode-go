package problems

/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	n := len(cost)
	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return dp[n-1]
}

// @lc code=end

// 1, 100, 1, 1, 1, 100, 1, 1, 100, 1
// 1 100 2

// dp[i] = cost[i] + min(dp[i-1], dp[i-2])
// 优化空间
func MinCostClimbingStairs2(cost []int) int {
	cost = append(cost, 0)
	n := len(cost)
	// dp := make([]int, n)
	dp_i_1 := cost[0]
	dp_i_2 := cost[1]
	ans := 0
	for i := 2; i < n; i++ {
		ans = cost[i] + min(dp_i_1, dp_i_2)
		dp_i_1, dp_i_2 = dp_i_2, ans
	}
	return ans
}
