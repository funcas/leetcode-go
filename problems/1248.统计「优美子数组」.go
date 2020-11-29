package problems

/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] 统计「优美子数组」
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	l := len(nums)
	a := make([]int, 1)
	a[0] = -1
	for i := 0; i < l; i++ {
		if nums[i]%2 != 0 {
			a = append(a, i)
		}
	}
	a = append(a, l)
	ret := 0
	for i := 1; i < len(a); i++ {
		if i+k < len(a) {
			ret += (a[i] - a[i-1]) * (a[i+k] - a[i+k-1])
		}
	}
	return ret
}

// @lc code=end
