package problems

import "sort"

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		if len(ans) == 0 || interval[0] > ans[len(ans)-1][1] {
			ans = append(ans, []int{interval[0], interval[1]})
		} else {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], interval[1])
		}
	}
	return ans
}

// @lc code=end
