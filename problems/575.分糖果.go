package problems

/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candyType []int) int {
	km := make(map[int]struct{})
	n := len(candyType)
	for _, v := range candyType {
		km[v] = struct{}{}
	}
	k := len(km)
	res := n / 2
	if res <= k {
		return res
	}
	return k

}

// @lc code=end
