package t2547

import "math"

// 抄的答案
func minCost(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n+1)
	for i := 0; i < n; i++ {
		state, unique, mn := make([]int8, n), 0, math.MaxInt
		for j := i; j >= 0; j-- {
			x := nums[j]
			if state[x] == 0 { // 首次出现
				state[x] = 1
				unique++
			} else if state[x] == 1 { // 不再唯一
				state[x] = 2
				unique--
			}
			mn = min(mn, f[j]-unique)
		}
		f[i+1] = mn + k
	}
	return f[n] + n
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
