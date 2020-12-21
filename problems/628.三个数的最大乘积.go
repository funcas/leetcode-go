package problems

import "sort"

/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	return max(nums[n-1]*nums[n-2]*nums[n-3], nums[n-1]*nums[0]*nums[1])
}

// @lc code=end
