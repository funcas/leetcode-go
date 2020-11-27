package problems

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	for lastZero, cur := 0, 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[cur], nums[lastZero] = nums[lastZero], nums[cur]
			lastZero++
		}
	}
}

// @lc code=end
