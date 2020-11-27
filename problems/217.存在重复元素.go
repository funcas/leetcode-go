package problems

/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	tmp := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := tmp[num]; exists {
			return true
		}
		tmp[num] = struct{}{}
	}
	return false
}

// @lc code=end
