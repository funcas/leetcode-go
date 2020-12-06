package problems

/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 */
// 分治，对两个数组按k组合拆分，分别求出合并出来的最大值，最后对每组最大值取最大值即为答案
// @lc code=start
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	ans := []int{}
	for i := 0; i < k+1; i++ {
		if i <= len(nums1) && k-i <= len(nums2) {
			tmp := merge(pick(nums1, i), pick(nums2, k-i))
			if compare(ans, tmp) {
				ans = tmp
			}
		}
	}
	return ans
}

// if a < b return true else return false
func compare(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func pick(nums []int, k int) []int {
	s := []int{}
	drop := len(nums) - k
	for _, v := range nums {
		for drop > 0 && len(s) > 0 && s[len(s)-1] < v {
			s = s[:len(s)-1]
			drop--
		}
		s = append(s, v)
	}
	return s[:k]
}

func merge(nums1 []int, nums2 []int) []int {
	n1 := len(nums1)
	n2 := len(nums2)
	ans := make([]int, n1+n2)

	for i := 0; len(nums1) > 0 || len(nums2) > 0; i++ {
		if compare(nums1, nums2) {
			ans[i], nums2 = nums2[0], nums2[1:]
		} else {
			ans[i], nums1 = nums1[0], nums1[1:]
		}
	}

	return ans
}

// @lc code=end
