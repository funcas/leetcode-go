package problems

import "fmt"

/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	ans := []string{}
	tmp := [][]int{}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			tmp = append(tmp, nums[j:i])
			j = i
		}
	}
	if j < len(nums) {
		tmp = append(tmp, nums[j:])
	}
	for _, v := range tmp {
		if len(v) == 1 {
			ans = append(ans, fmt.Sprintf("%d", v[0]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", v[0], v[len(v)-1]))
		}
	}
	return ans
}

// @lc code=end
