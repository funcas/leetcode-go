package problems

/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	n := len(A) - 1 // 必为奇数
	i, j := 0, 1    // 奇偶指针

	for j <= n && i <= n-1 {
		for i <= n && A[i]%2 == 0 {
			i += 2
		}
		for j <= n-1 && A[j]%2 != 0 {
			j += 2
		}
		if i <= n && j <= n {
			A[i], A[j] = A[j], A[i]
		}
		i += 2
		j += 2
	}

	return A
}

// @lc code=end
