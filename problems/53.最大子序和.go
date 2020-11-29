package problems

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	p, m := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		p = max(p+nums[i], nums[i])
		m = max(m, p)
	}
	return m
}

// @lc code=end
