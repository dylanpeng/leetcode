package result

/*
115. 不同的子序列

给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。

字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）

题目数据保证答案符合 32 位带符号整数范围。

示例 1：

输入：s = "rabbbit", t = "rabbit"
输出：3
解释：
如下图所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
(上箭头符号 ^ 表示选取的字母)
rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^

示例 2：

输入：s = "babgbag", t = "bag"
输出：5
解释：
如下图所示, 有 5 种可以从 s 中得到 "bag" 的方案。
(上箭头符号 ^ 表示选取的字母)
babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
    ^^^

提示：

    0 <= s.length, t.length <= 1000
    s 和 t 由英文字母组成

*/

// 思路：部分穷举，计数，leetcode时间太长，不过
func NumDistinct(s string, t string) int {
	sRune, tRune := GetRune(s, t)
	tLen, i, count := len(t), 0, 0

	if tLen == 1 {
		for ; i < len(s); i++ {
			if s[i] == t[0] {
				count++
			}
		}
		return count
	}

	tempList := make([][]rune, 0)

	for ; i < len(sRune); i++ {
		newList := make([][]rune, 0)

		for _, str := range tempList {
			if len(str) < tLen {
				newList = append(newList, str)
			} else {
				continue
			}

			newStr := make([]rune, 0)
			newStr = append(newStr, str...)
			newStr = append(newStr, sRune[i])

			if len(newStr) < tLen {
				if IsSame(newStr, tRune[:len(newStr)]) {
					newList = append(newList, newStr)
				}
			} else if len(newStr) == tLen && IsSame(newStr, tRune) {
				count++
			}
		}

		if sRune[i] == tRune[0] {
			newList = append(newList, []rune{sRune[i]})
		}

		tempList = newList
	}

	return count
}

func IsSame(a []rune, b []rune) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func GetRune(s string, t string) (sRune []rune, tRune []rune) {
	tRune, sRune = []rune(t), make([]rune, 0)
	sR := []rune(s)
	tMap := make(map[rune]bool)

	for i := 0; i < len(t); i++ {
		tMap[tRune[i]] = true
	}

	for i := 0; i < len(s); i++ {
		if tMap[sR[i]] {
			sRune = append(sRune, sR[i])
		}
	}

	return
}

// 思路：tCount记录到当前位置拼成功的子序列的个数，如t:"abcd" tCount[0]:记录拼成"a"的个数 tCount[2]:记录拼成"abc"的个数
// 每当与当前字符相等时，当前字符拼成个数为上一位拼成个数加上未拼时的当前个数
func NumDistinct2(s string, t string) int {
	sRune, tRune := GetRune(s, t)
	tLen, i, j := len(t), 0, 0

	tCount := make([]int, tLen)

	for i = 0; i < len(sRune); i++ {
		for j = tLen - 1; j >= 0; j-- {
			if sRune[i] == tRune[j] {
				if j == 0 {
					tCount[0]++
				} else if tCount[j-1] > 0 {
					tCount[j] += tCount[j-1]
				}
			}
		}
	}

	return tCount[tLen-1]
}

func NumDistinct3(s string, t string) int {
	tLen, i, j := len(t), 0, 0

	tCount := make([]int, tLen)

	for i = 0; i < len(s); i++ {
		for j = tLen - 1; j >= 0; j-- {
			if s[i] == t[j] {
				if j == 0 {
					tCount[0]++
				} else if tCount[j-1] > 0 {
					tCount[j] += tCount[j-1]
				}
			}
		}
	}

	return tCount[tLen-1]
}
