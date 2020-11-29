package problems

/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] 数组中的最长山脉
 */

// @lc code=start
func longestMountain(A []int) int {
	isM := false
	count := 0
	j := 0
	for i := 1; i < len(A)-1; i++ {
		//down
		if A[i-1] > A[i] && A[i] <= A[i+1] {
			if isM {
				tmp := i - j + 1
				if count < tmp {
					count = tmp
				}
			}
			isM = false
			j = i
		}
		// up
		if A[i-1] < A[i] && A[i] > A[i+1] {
			isM = true
		}
		// last two
		if i == len(A)-2 && isM && A[i+1] < A[i] {
			tmp := i - j + 2
			if count < tmp {
				count = tmp
			}
		}
		if A[i] == A[i-1] {
			j = i
		}

	}
	return count
}

// @lc code=end
