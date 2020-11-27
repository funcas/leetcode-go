package problems

import "sort"

/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

// @lc code=start
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)

	ans := 0
	for i, j := 0, 1; j < len(nums); i, j = i+1, j+1 {
		tmp := nums[j] - nums[i]
		if tmp > ans {
			ans = tmp
		}
	}
	return ans
}

// @lc code=end
