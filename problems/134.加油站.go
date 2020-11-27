package problems

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	size := len(gas)
	for i := 0; i < size; i++ {
		remain := 0
		for j := i; j < size+i; j++ {
			idx := j
			if idx >= size {
				idx = idx - size
			}
			remain = (remain + gas[idx]) - cost[idx]
			if remain < 0 {
				break
			}
		}
		if remain >= 0 {
			return i
		}
	}
	return -1
}

// @lc code=end
