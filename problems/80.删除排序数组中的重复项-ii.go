package problems

/*
 * @lc app=leetcode.cn id=80 lang=golang
 *
 * [80] 删除排序数组中的重复项 II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	cur := nums[0]
	cnt := 1
	x := 10001
	for i := 1; i < len(nums); i++ {
		if nums[i] == cur {
			cnt++
		} else {
			cnt = 1
			cur = nums[i]
		}
		if cnt > 2 {
			nums[i] = x
			n--
		}

	}
	p := 0
	q := 0
	for p < n && nums[p] != x {
		p++
	}
	q = p
	for i := p; i < n; i++ {

		for q < len(nums) && nums[q] == x {
			q++
		}
		nums[i] = nums[q]
		q++
	}
	return n
}

// @lc code=end
