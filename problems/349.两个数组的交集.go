package problems

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]struct{})
	m2 := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}

	ans := []int{}
	for _, v := range nums2 {
		if _, ok := m2[v]; !ok {
			m2[v] = struct{}{}
			if _, exists := m[v]; exists {
				ans = append(ans, v)
			}
		}

	}
	return ans
}

// @lc code=end
