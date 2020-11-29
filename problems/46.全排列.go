package problems

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	var res [][]int
	perm(&res, nums, 0)
	return res
}

func perm(res *[][]int, nums []int, index int) {
	if index >= len(nums) {
		x := make([]int, len(nums))
		copy(x, nums)
		*res = append(*res, x)
	}
	for i := index; i < len(nums); i++ {
		swap(&nums[i], &nums[index])
		perm(res, nums, index+1)
		swap(&nums[index], &nums[i])
	}
}

// @lc code=end
