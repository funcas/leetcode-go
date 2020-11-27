package problems

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	m[nums[0]] = 0
	result := make([]int, 2)
	var exists bool
	val, o, l := 0, 0, len(nums)

	for i := 1; i < l; i++ {
		o = nums[i]
		if val, exists = m[target-o]; exists {
			result[0] = val
			result[1] = i
		}
		m[o] = i
	}

	return result
}

// @lc code=end
