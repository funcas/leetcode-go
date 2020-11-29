package problems

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	size := len(nums)
	if size == 1 {
		return true
	}
	jump := 0
	for i := 0; i < size-1; i++ {
		if i <= jump {
			jump = max(nums[i]+i, jump)
		}

	}
	return jump >= size-1
}

// @lc code=end
