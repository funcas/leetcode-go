package problems

/*
 * @lc app=leetcode.cn id=659 lang=golang
 *
 * [659] 分割数组为连续子序列
 */

// @lc code=start
func isPossible(nums []int) bool {
	n := len(nums)
	i, p1, p2, p3 := 0, 0, 0, 0

	for i < n {
		start := i
		x := nums[i]
		for i < n && nums[i] == x {
			i++
		}
		cnt := i - start
		if start > 0 && x != nums[start-1]+1 {
			if p1+p2 > 0 {
				return false
			}
			p1 = cnt
			p2, p3 = 0, 0
		} else {
			if p1+p2 > cnt {
				return false
			}
			remain := cnt - p1 - p2
			k := min(p3, remain)
			p3 = k + p2
			p2 = p1
			p1 = remain - k
		}
	}
	return p1 == 0 && p2 == 0

}

// @lc code=end
