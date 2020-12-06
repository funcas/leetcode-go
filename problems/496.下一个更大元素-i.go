package problems

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		j := i + 1
		for j < len(nums2) && nums2[j]-nums2[i] < 0 {
			j++
		}
		if j == len(nums2) {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = nums2[j]
		}

	}
	ans := []int{}
	for _, v := range nums1 {
		ans = append(ans, m[v])
	}
	return ans

}

// @lc code=end
