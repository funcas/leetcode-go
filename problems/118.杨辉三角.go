package problems

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	ans := [][]int{}
	for i := 0; i < numRows; i++ {
		tmp := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, ans[i-1][j]+ans[i-1][j-1])
			}
		}
		ans = append(ans, tmp)
	}
	return ans
}

// @lc code=end
