package problems

import "sort"

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	if len(candidates) == 0 {
		return ans
	}
	var backtrace func(idx int, retain int, path []int)
	backtrace = func(idx int, retain int, path []int) {

		if retain == 0 {
			// find it
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := idx; i < len(candidates); i++ {
			if retain-candidates[i] < 0 {
				break
			}
			if i > idx && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrace(i+1, retain-candidates[i], path)
			path = path[:len(path)-1]
		}

	}
	sort.Ints(candidates)
	backtrace(0, target, []int{})
	return ans
}

// @lc code=end
