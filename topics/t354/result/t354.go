package result

import (
	"fmt"
	"sort"
)

/*
354. 俄罗斯套娃信封问题

给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:

输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。


*/

type sortEnvelopes [][]int

func (s sortEnvelopes) Len() int {
	return len(s)
}
func (s sortEnvelopes) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] > s[j][1]
	}
	return s[i][0] < s[j][0]
}
func (s sortEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func MaxEnvelopes(envelopes [][]int) int {
	sort.Sort(sortEnvelopes(envelopes))
	fmt.Printf("%+v", envelopes)
	dp := []int{}
	for _, e := range envelopes {
		low, high := 0, len(dp)
		for low < high {
			mid := low + (high-low)>>1
			if dp[mid] >= e[1] {
				high = mid
			} else {
				low = mid + 1
			}
		}
		if low == len(dp) {
			dp = append(dp, e[1])
		} else {
			dp[low] = e[1]
		}
	}
	return len(dp)
}
