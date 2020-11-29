package problems

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	primes := make([]bool, n)
	for i := 0; i < n; i++ {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false
	sq := sqrt1(n)
	for i := 2; i <= sq; i++ {
		if !primes[i] {
			continue
		}
		for multi := i << 1; multi < n; multi += i {
			primes[multi] = false
		}
	}
	count := 0
	for i := 0; i < len(primes); i++ {
		if primes[i] {
			count++
		}
	}
	return count
}
func sqrt1(x int) int {
	if x == 0 {
		return 0
	}
	var last float64 = 0
	var res float64 = 1
	for res != last {
		last = res
		res = (res + float64(x)/res) / 2
	}
	return int(res)
}

// @lc code=end
