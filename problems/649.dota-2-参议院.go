package problems

/*
 * @lc app=leetcode.cn id=649 lang=golang
 *
 * [649] Dota2 参议院
 */

// @lc code=start
func PredictPartyVictory(senate string) string {
	rad := []int{}
	dire := []int{}
	n := len(senate)
	for i := 0; i < n; i++ {
		s := senate[i]
		if s == 'R' {
			rad = append(rad, i)
		} else {
			dire = append(dire, i)
		}
	}
	for len(rad) != 0 && len(dire) != 0 {
		r := rad[0]
		d := dire[0]
		rad = rad[1:]
		dire = dire[1:]
		if r < d {
			rad = append(rad, n+r)
		} else {
			dire = append(dire, n+d)
		}

	}
	if len(rad) == 0 {
		return "Dire"
	} else {
		return "Radiant"
	}

}

// @lc code=end
