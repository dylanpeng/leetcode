package result

/*
1178. 猜字谜

外国友人仿照中国字谜设计了一个英文版猜字谜小游戏，请你来猜猜看吧。

字谜的迷面 puzzle 按字符串形式给出，如果一个单词 word 符合下面两个条件，那么它就可以算作谜底：

    单词 word 中包含谜面 puzzle 的第一个字母。
    单词 word 中的每一个字母都可以在谜面 puzzle 中找到。
    例如，如果字谜的谜面是 "abcdefg"，那么可以作为谜底的单词有 "faced", "cabbage", 和 "baggage"；而 "beefed"（不含字母 "a"）以及 "based"（其中的 "s" 没有出现在谜面中）。

返回一个答案数组 answer，数组中的每个元素 answer[i] 是在给出的单词列表 words 中可以作为字谜迷面 puzzles[i] 所对应的谜底的单词数目。



示例：

输入：
words = ["aaaa","asas","able","ability","actt","actor","access"],
puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
输出：[1,1,3,2,4,0]
解释：
1 个单词可以作为 "aboveyz" 的谜底 : "aaaa"
1 个单词可以作为 "abrodyz" 的谜底 : "aaaa"
3 个单词可以作为 "abslute" 的谜底 : "aaaa", "asas", "able"
2 个单词可以作为 "absoryz" 的谜底 : "aaaa", "asas"
4 个单词可以作为 "actresz" 的谜底 : "aaaa", "asas", "actt", "access"
没有单词可以作为 "gaswxyz" 的谜底，因为列表中的单词都不含字母 'g'。



提示：

    1 <= words.length <= 10^5
    4 <= words[i].length <= 50
    1 <= puzzles.length <= 10^4
    puzzles[i].length == 7
    words[i][j], puzzles[i][j] 都是小写英文字母。
    每个 puzzles[i] 所包含的字符都不重复。


*/
func FindNumOfValidWords(words []string, puzzles []string) []int {
	wLen, pLen, i, j, num, bitNum, r, k, v, ok, key := len(words), len(puzzles), 0, 0, 0, 0, make([]rune, 0), 0, 0, false, 0
	result := make([]int, pLen)
	wordsMap := make(map[int]int)
	puzzlesMap := make(map[int]int)

	for i = 0; i < wLen; i++ {
		num = 0
		r = []rune(words[i])
		for j = 0; j < len(r); j++ {
			bitNum = 1 << (r[j] - 97)
			if num&bitNum != bitNum {
				num += bitNum
			}
		}

		if v, ok = wordsMap[num]; ok {
			wordsMap[num] = v + 1
		} else {
			wordsMap[num] = 1
		}
	}

	for i = 0; i < pLen; i++ {
		num = 0
		r = []rune(puzzles[i])
		num = (1 << (r[0] - 97)) + (1 << (r[1] - 97)) + (1 << (r[2] - 97)) + (1 << (r[3] - 97)) + (1 << (r[4] - 97)) + (1 << (r[5] - 97)) + (1 << (r[6] - 97))
		bitNum = 1 << (r[0] - 97)

		key = int(r[0]<<26) + num
		if v, ok = puzzlesMap[key]; ok {
			result[i] = v
			continue
		}

		for k, v = range wordsMap {
			if bitNum&k != bitNum || k&num != k {
				continue
			}

			result[i] += v
		}

		puzzlesMap[key] = result[i]
	}

	return result
}

/*
思路：采用位运算，26位代表26个字母，每个word算出对应值的个数放入map中，遍历puzzle的子集算出总数
*/
func FindNumOfValidWords2(words []string, puzzles []string) []int {
	count := make(map[int]int, len(words))
	for _, w := range words {
		count[mask(w)]++
	}

	res := make([]int, len(puzzles))

	for i, p := range puzzles {
		subs := subsWithHead(p)
		for _, s := range subs {
			res[i] += count[s]
		}
	}

	return res
}

func mask(w string) int {
	res := 0
	for _, l := range w {
		res |= 1 << uint(l-'a')
	}
	return res
}

func subsWithHead(p string) []int {
	n := len(p)
	res := make([]int, 1, 128)
	res[0] = 1 << uint(p[0]-'a')
	for i := 1; i < n; i++ {
		x := 1 << uint(p[i]-'a')
		for _, r := range res {
			res = append(res, r|x)
		}
	}
	return res
}
