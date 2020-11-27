package problems

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	i, j := 0, len(nums)
	pos := -1
	ans := []int{-1, -1}
	if len(nums) == 0 || target > nums[len(nums)-1] {
		return ans
	}
	for i <= j {
		mid := (i + j) / 2
		if target == nums[mid] {
			pos = mid
			break
		}
		if target < nums[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	if pos >= 0 {
		m, n := pos, pos
		for m > 0 {
			if m >= 1 && nums[m] != nums[m-1] {
				break
			}
			m--
		}
		for n < len(nums)-1 {
			if nums[n] != nums[n+1] {
				break
			}
			n++
		}
		ans = []int{m, n}
	}

	return ans
}

// @lc code=end
