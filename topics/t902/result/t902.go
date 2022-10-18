package result

import (
	"strconv"
)

/*
902. 最大为 N 的数字组合
给定一个按 非递减顺序 排列的数字数组 digits 。你可以用任意次数 digits[i] 来写的数字。例如，如果 digits = ['1','3','5']，我们可以写数字，如 '13', '551', 和 '1351315'。

返回 可以生成的小于或等于给定整数 n 的正整数的个数 。



示例 1：

输入：digits = ["1","3","5","7"], n = 100
输出：20
解释：
可写出的 20 个数字是：
1, 3, 5, 7, 11, 13, 15, 17, 31, 33, 35, 37, 51, 53, 55, 57, 71, 73, 75, 77.
示例 2：

输入：digits = ["1","4","9"], n = 1000000000
输出：29523
解释：
我们可以写 3 个一位数字，9 个两位数字，27 个三位数字，
81 个四位数字，243 个五位数字，729 个六位数字，
2187 个七位数字，6561 个八位数字和 19683 个九位数字。
总共，可以使用D中的数字写出 29523 个整数。
示例 3:

输入：digits = ["7"], n = 8
输出：1


提示：

1 <= digits.length <= 9
digits[i].length == 1
digits[i] 是从 '1' 到 '9' 的数
digits 中的所有值都 不同
digits 按 非递减顺序 排列
1 <= n <= 109
*/

func AtMostNGivenDigitSet(digits []string, n int) int {
	list, listLen := Sort902(digits)
	nList, nLen := LenInt(n)

	count := Count902(list, nList)

	if nLen > 1 {
		for i := 1; i < nLen; i++ {
			count += Pow902(listLen, i)
		}
	}

	return count
}

func LenInt(n int) (nList []int, l int) {
	nStr := strconv.Itoa(n)
	nList = make([]int, 0, 8)

	for i := 0; i < len(nStr); i++ {
		dN, _ := strconv.Atoi(string(nStr[i]))
		nList = append(nList, dN)
		l++
	}

	return
}

func Sort902(digits []string) (result []int, l int) {
	temp := map[string]int{}
	result = make([]int, 0, 8)

	for _, d := range digits {
		if _, exit := temp[d]; !exit {
			temp[d] = 1
			dN, _ := strconv.Atoi(d)
			result = append(result, dN)
			l++
		}
	}

	return
}

func Count902(list, nList []int) (result int) {
	x := len(list)
	curNum := nList[0]
	leftLen := len(nList) - 1

	if leftLen < 0 {
		return 0
	}

	for _, d := range list {
		if d < curNum {
			result += Pow902(x, leftLen)
		} else if d == curNum {
			if leftLen == 0 {
				return result + 1
			}
			return result + Count902(list, nList[1:])
		} else {
			return
		}
	}

	return
}

func Pow902(x, y int) int {
	if y == 0 {
		return 1
	}

	if y < 0 {
		return 0
	}

	result := x
	for i := 1; i < y; i++ {
		result *= x
	}

	return result
}
