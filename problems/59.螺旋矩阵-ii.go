package problems

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	c := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for left <= right && top <= bottom {

		for i := left; i <= right; i++ {
			ans[top][i] = c
			c++
		}
		for i := top + 1; i <= bottom; i++ {
			ans[i][right] = c
			c++
		}
		for i := right - 1; i >= left; i-- {
			ans[bottom][i] = c
			c++
		}
		for i := bottom - 1; i > top; i-- {
			ans[i][left] = c
			c++
		}
		left++
		right--
		top++
		bottom--
	}
	return ans
}

// @lc code=end

/*

[1 2 3 4 5]
[16 17 18 19 6]
[15 24 25 20 7]
[14 23 22 21 8]
[13 12 11 10 9]
*/
