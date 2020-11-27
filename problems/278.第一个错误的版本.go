package problems

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	i, j := 1, n

	for i <= j {
		mid := (i + j) / 2
		if isBadVersion(mid) {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i
}

func isBadVersion(version int) bool {
	if version >= 2 {
		return true
	}
	return false
}

// @lc code=end
