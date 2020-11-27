package problems

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 1
	}
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := range nums {
		idx := abs(nums[i])
		if idx <= n {
			nums[idx-1] = -abs(nums[idx-1])
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

// @lc code=end
