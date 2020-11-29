package problems

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	res := make([]int, l+1)
	res[0] = 0
	res[1] = nums[0]
	for i := 2; i <= l; i++ {
		res[i] = max(res[i-1], res[i-2]+nums[i-1])
	}
	return res[l]
}

// @lc code=end
