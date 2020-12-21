package problems

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func FindLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	i, j := 0, 1
	ans := 0

	for i < n && j < n {
		if nums[j-1] >= nums[j] {
			tmp := j - i
			if ans < tmp {
				ans = tmp
			}
			i = j
		}
		j++

	}
	tmp := j - i
	if ans < tmp {
		ans = tmp
	}
	return ans
}

// @lc code=end
/*

1 3 5 4 7


*/
