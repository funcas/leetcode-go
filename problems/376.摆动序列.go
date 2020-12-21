package problems


/*
 * @lc app=leetcode.cn id=376 lang=golang
 *
 * [376] 摆动序列
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up, down := 1, 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up = down + 1
		}
		if nums[i] < nums[i-1] {
			down = up + 1
		}
	}
	return max(up, down)
}

// @lc code=end

//  1 17 5 10 13 15 10 5 16 8

/*
   17     10 13 15        16
1      5            10 5      8
*/
