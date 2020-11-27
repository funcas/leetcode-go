package problems

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	lmax := make([]int, n)
	rmax := make([]int, n)
	lmax[0] = height[0]
	rmax[n-1] = height[n-1]
	ans := 0
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		ans += min(lmax[i], rmax[i]) - height[i]
	}

	return ans
}

// @lc code=end
// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
