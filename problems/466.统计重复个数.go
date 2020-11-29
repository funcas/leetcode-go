package problems

/*
 * @lc app=leetcode.cn id=466 lang=golang
 *
 * [466] 统计重复个数
 */

// @lc code=start
type sTuple struct {
	S1Cnt int
	S2Cnt int
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	m := make(map[int]*sTuple)
	s1Count, s2Count, idx := 0, 0, 0
	var unLoop *sTuple
	var loop *sTuple
	for {
		s1Count++
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[idx] {
				idx++
				if idx == len(s2) {
					s2Count++
					idx = 0
				}
			}
		}
		// s1遍历结束了
		if n1 == s1Count {
			return s2Count / n2
		}

		if _, ok := m[idx]; !ok {
			m[idx] = &sTuple{
				S1Cnt: s1Count,
				S2Cnt: s2Count,
			}
		} else { // 找到了循环点
			unLoop = m[idx]
			loop = &sTuple{
				S1Cnt: s1Count - unLoop.S1Cnt,
				S2Cnt: s2Count - unLoop.S2Cnt,
			}
			break
		}
	}
	// 已知的元素数量
	ret := unLoop.S2Cnt + (n1-unLoop.S1Cnt)/loop.S1Cnt*loop.S2Cnt
	// 剩下多少个元素
	remain := (n1 - unLoop.S1Cnt) % loop.S1Cnt
	for i := 0; i < remain; i++ {
		for j := 0; j < len(s1); j++ {
			if s1[j] == s2[idx] {
				idx++
				if idx == len(s2) {
					ret++
					idx = 0
				}
			}
		}
	}
	return ret / n2
}

// @lc code=end
