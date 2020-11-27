package problems

/*
 * @lc app=leetcode.cn id=514 lang=golang
 *
 * [514] 自由之路
 */

// @lc code=start
func findRotateSteps(ring string, key string) int {
	lr, lk := len(ring), len(key)
	pos := make(map[uint8][]int)
	for i := 0; i < len(ring); i++ {
		pos[ring[i]] = append(pos[ring[i]], i)
	}
	dp := make([][]int, lk)
	for i := 0; i < lk; i++ {
		dp[i] = make([]int, lr)
		for j := 0; j < lr; j++ {
			dp[i][j] = 1<<31 - 1
		}
	}
	for _, p := range pos[key[0]] {
		dp[0][p] = min2(p, lr-p) + 1
	}
	for i := 1; i < lk; i++ {
		for _, j := range pos[key[i]] {
			for _, k := range pos[key[i-1]] {
				dp[i][j] = min2(dp[i][j], dp[i-1][k]+min2(abs(j-k), lr-abs(j-k))+1)
			}
		}
	}

	return min2(dp[lk-1]...)
}

func min2(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

// @lc code=end
