package problems

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		if v, ok := m[nums[i]]; ok {
			m[nums[i]] = v + 1
		} else {
			m[nums[i]] = 1
		}
	}
	max := 0
	maxI := 0
	for k, v := range m {
		if max < v {
			max = v
			maxI = k
		}
	}
	return maxI
}

// @lc code=end
