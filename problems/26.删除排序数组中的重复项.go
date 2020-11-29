package problems

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates_(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	current := nums[0]
	i := 0
	for j := 1; j < l; j++ {
		if nums[i] != nums[j] {
			i++
			current = nums[j]
			nums[i] = current
		}

	}
	return i + 1
}

// @lc code=end
