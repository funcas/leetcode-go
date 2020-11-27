package problems

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	wordMap := make(map[string]struct{})
	for _, s := range wordList {
		if s != beginWord {
			wordMap[s] = struct{}{}
		}
	}
	if _, ok := wordMap[endWord]; !ok {
		return 0
	}
	q := []string{beginWord}
	visited := make(map[string]struct{})
	visited[beginWord] = struct{}{}
	var changeWord func(beginWord string, endWord string) (r bool, tq []string)
	step := 1

	changeWord = func(beginWord string, endWord string) (r bool, tq []string) {
		var tmp []string
		beginWordByte := []byte(beginWord)
		for i := 0; i < len(beginWordByte); i++ {
			ori := beginWordByte[i]
			for j := 'a'; j <= 'z'; j++ {
				beginWordByte[i] = byte(j)
				nextWord := string(beginWordByte)
				if _, ok := wordMap[nextWord]; ok {
					if nextWord == endWord {
						return true, tmp
					}
					if _, ok := visited[nextWord]; !ok {
						tmp = append(tmp, nextWord)
						visited[nextWord] = struct{}{}
					}
				}
			}
			beginWordByte[i] = ori
		}
		return false, tmp
	}
	for len(q) > 0 {
		var tmp []string
		for _, item := range q {
			ok, qq := changeWord(item, endWord)
			if ok {
				return step + 1
			}
			tmp = append(tmp, qq...)
		}
		q = tmp
		step++
	}
	return 0
}

// @lc code=end
