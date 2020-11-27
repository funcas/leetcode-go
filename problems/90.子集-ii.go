package problems

import "sort"

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	size := len(nums)
	ans := [][]int{}
	sort.Ints(nums)
	var dfs func(idx int, path []int)
	dfs = func(idx int, path []int) {
		if idx > size {
			return
		}
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		for i := idx; i < size; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{})
	return ans
}

// @lc code=end
