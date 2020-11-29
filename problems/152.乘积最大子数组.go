package problems

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := ^(int(^uint(0) >> 1))
	for i := 0; i < len(nums); i++ {
		c := 1
		for j := i; j < len(nums); j++ {
			c = c * nums[j]
			if max < c {
				max = c
			}
		}

	}
	return max
}

// @lc code=end
