package problems

import "sort"

/*
 * @lc app=leetcode.cn id=1365 lang=golang
 *
 * [1365] 有多少小于当前数字的数字
 */

// @lc code=start
func smallerNumbersThanCurrent(nums []int) []int {
	n := make([]int, len(nums))
	for i, v := range nums {
		n[i] = v
	}
	sort.Ints(n)
	tmp := make(map[int]int)
	for idx, v := range n {
		if _, ok := tmp[v]; !ok {
			tmp[v] = idx
		}
	}
	ans := make([]int, len(nums))
	for i, v := range nums {
		if i < 0 && v == nums[i-1] {
			ans[i] = ans[i-1]
		} else {
			ans[i] = tmp[v]
		}
	}
	return ans
}

// @lc code=end
