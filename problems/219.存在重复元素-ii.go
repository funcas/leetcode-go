package problems

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if r, ok := m[v]; ok {
			if i-r <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}

// @lc code=end
