package problems

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */
// sum(nums) - min(nums)*len(nums)
// @lc code=start
func minMoves(nums []int) int {
	n := len(nums)
	min := 1<<63 - 1
	sum := 0
	for _, v := range nums {
		if v < min {
			min = v
		}
		sum += v
	}
	return sum - min*n

}

// @lc code=end
