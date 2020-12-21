package problems

import "sort"

/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 */

// @lc code=start
func ArrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := len(nums) - 1; i >= 0; i -= 2 {
		sum += nums[i-1]
	}
	return sum
}

// @lc code=end
