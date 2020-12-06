package problems

/*
 * @lc app=leetcode.cn id=65 lang=golang
 *
 * [65] 有效数字
 */

// @lc code=start
func IsNumber(s string) bool {
	state := 0
	finals := []int{0, 0, 0, 1, 0, 1, 1, 0, 1}
	transfer := [][]int{{0, 1, 6, 2, -1},
		{-1, -1, 6, 2, -1},
		{-1, -1, 3, -1, -1},
		{8, -1, 3, -1, 4},
		{-1, 7, 5, -1, -1},
		{8, -1, 5, -1, -1},
		{8, -1, 6, 3, 4},
		{-1, -1, 5, -1, -1},
		{8, -1, -1, -1, -1}}
	for i := 0; i < len(s); i++ {
		st := trans(s[i])
		if st < 0 {
			return false
		}
		state = transfer[state][st]
		if state < 0 {
			return false
		}
	}
	return finals[state] == 1

}

func trans(c uint8) int {

	switch c {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '.':
		return 3
	case 'e':
		return 4
	default:
		if c >= '0' && c <= '9' {
			return 2
		}
	}
	return -1
}

// @lc code=end

/**
state | blank  | +/-  |  0-9  |  .  |  e  |  other  |
 0	  | 0 	   |1	  |6      |2	 -1		-1
 1    | -1	   |-1	  |6      |2 	 -1		-1
 2    | -1	   |-1	  | 3     |-1	 -1		-1
 3    | 8	   |7	  |3      |-1 	 4		-1
 4    | -1	   |-1	  |5      |-1 	 -1		-1
 5    |  8	   |-1	  |5      |-1	 -1		-1
 6    |  8	   |-1	  |6      |3	 4		-1
 7    |  -1	   |-1	  |5      |-1	 -1		-1
 8    |  8	   |-1	  | -1    |-1	 -1		-1
*/
