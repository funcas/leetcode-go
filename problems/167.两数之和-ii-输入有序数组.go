package problems

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */
// 1 2 4 4 5 6 | 3
// @lc code=start
func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		res := target - numbers[j]
		if res == numbers[i] {
			return []int{i + 1, j + 1}
		}
		if res < numbers[i] {
			j--
		} else {
			i++
		}
	}
	return []int{}
}

// @lc code=end
