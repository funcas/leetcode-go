package problems

/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	flag := 0
	if A[1] < A[0] {
		return false
	}
	for i := 1; i < len(A); i++ {
		if flag == 0 && A[i] == A[i-1] {
			return false
		}
		if A[i] < A[i-1] && flag == 0 {
			flag = 1
		}
		if flag == 1 && A[i] >= A[i-1] {
			return false
		}
	}
	return flag == 1
}

func validMountainArray2(A []int) bool {
	n := len(A)
	i := 0

	for i+1 < n && A[i+1] > A[i] {
		i++
	}

	if i == 0 || i == n-1 {
		return false
	}

	for i+1 < n && A[i+1] < A[i] {
		i++
	}
	return i == n-1

}

// @lc code=end
