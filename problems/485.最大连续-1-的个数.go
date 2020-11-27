package problems

/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	ans := 0
	tmp := 0
	for i := range nums {
		if nums[i]&1 == 1 {
			tmp++
		} else {
			if ans < tmp {
				ans = tmp
			}
			tmp = 0
		}
	}
	if tmp > 0 && ans < tmp {
		ans = tmp
	}
	return ans
}

// @lc code=end
