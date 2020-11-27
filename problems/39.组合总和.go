package problems

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var backtrace func(idx int, sum int, path []int)
	ans := [][]int{}
	backtrace = func(idx int, sum int, path []int) {
		if sum > target || idx == len(candidates) {
			return
		}
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		backtrace(idx, sum+candidates[idx], append(path, candidates[idx]))
		backtrace(idx+1, sum, append(path))

	}

	backtrace(0, 0, []int{})
	return ans
}

// @lc code=end
