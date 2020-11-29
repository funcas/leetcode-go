package problems

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */
// https://funcas.top/2020/04/18/post/
// @lc code=start
func maxArea(height []int) int {
	length := len(height) - 1
	maxArea := min(height[0], height[length]) * length
	for i, j := 0, length; i < j; {
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
		maxArea = max(maxArea, min(height[i], height[j])*(j-i))
	}
	return maxArea
}

// @lc code=end
