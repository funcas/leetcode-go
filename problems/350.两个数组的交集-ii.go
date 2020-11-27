package problems

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	ans := []int{}
	m := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range nums2 {
		if x, ok := m[v]; ok && x > 0 {
			ans = append(ans, v)
			m[v]--
		}
	}

	return ans
}

// @lc code=end
