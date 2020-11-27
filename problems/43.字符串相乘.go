package problems

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := make([]uint8, len(num1))
	n2 := make([]uint8, len(num2))
	n3 := make([]uint8, 0)
	for i, j := len(num1)-1, 0; i >= 0; i, j = i-1, j+1 {
		n1[j] = num1[i] - '0'
	}
	for i, j := len(num2)-1, 0; i >= 0; i, j = i-1, j+1 {
		n2[j] = num2[i] - '0'
	}
	flag := uint8(0)
	k := 0
	for i := 0; i < len(num2); i++ {
		k = i
		for j := 0; j < len(num1); j++ {
			if len(n3) <= k {
				n3 = append(n3, 0)
			}
			tmp := n2[i]*n1[j] + n3[k] + flag
			flag = tmp / 10
			n3[k] = tmp % 10
			k++
		}
		if flag > 0 {
			n3 = append(n3, flag)
			flag = 0
		}
	}

	for i, j := 0, len(n3)-1; i <= j; i, j = i+1, j-1 {
		if i != j {
			n3[i], n3[j] = n3[j], n3[i]
			n3[i] += '0'
			n3[j] += '0'
		} else {
			n3[i] += '0'
		}

	}

	return string(n3)
}

// @lc code=end
