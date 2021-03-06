package result

/*
132. 分割回文串 II

给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。

返回符合要求的 最少分割次数 。



示例 1：

输入：s = "aab"
输出：1
解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。

示例 2：

输入：s = "a"
输出：0

示例 3：

输入：s = "ab"
输出：1



提示：

    1 <= s.length <= 2000
    s 仅由小写英文字母组成

*/

func MinCut(s string) int {
	sArry, i, j, count, min := []rune(s), 0, 0, 0, 0

	result := make([]int, len(s))
	result[0] = 0
	for i = 1; i < len(s); i++ {
		count, min = 0, result[i-1]+1
		for j = 0; j <= i; j++ {
			if sArry[i] == sArry[j] {
				if IsCircle(s[j : i+1]) {
					if j == 0 {
						count = 0
					} else {
						count = result[j-1] + 1
					}
				} else {
					count = result[i-1] + 1
				}

				if count < min {
					min = count
				}
			}
		}
		result[i] = min
	}

	return result[len(s)-1]
}

func IsCircle(s string) bool {
	sArry, i, j := []rune(s), 0, len(s)-1

	for i <= j {
		if sArry[i] != sArry[j] {
			return false
		}
		i++
		j--
	}

	return true
}
