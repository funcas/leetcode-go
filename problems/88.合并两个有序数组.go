package problems

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge88(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
	}
	if n == 0 {
		return
	}
	var tmp []int
	a, b, i := 0, 0, 0

	for a < m && b < n {
		if nums1[a] < nums2[b] {
			tmp = append(tmp, nums1[a])
			i++
			a++
		} else {
			tmp = append(tmp, nums2[b])
			i++
			b++
		}
	}
	if a < m {
		tmp = append(tmp, nums1[a:]...)
	}
	if b < n {
		tmp = append(tmp, nums2[b:]...)
	}
	copy(nums1, tmp)
}

// @lc code=end
