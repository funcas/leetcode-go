package problems

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */
/// 1. 从右向左找到nums[i] < nums[i+1],得到较小数
/// 2. 从右向左找到nums[j] > nums[i]， 得到较大数
/// 3. 交换nums[i] nums[j]
/// 4. 将i以后的序列倒序
// @lc code=start
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j > i && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1)

}

func reverse(nums []int, begin int) {
	for i, j := begin, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end
