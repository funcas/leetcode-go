package problems

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i]) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

// func abs(num int) int {
// 	if num < 0 {
// 		return -num
// 	}
// 	return num
// }

// @lc code=end
