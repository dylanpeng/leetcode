package result

/*
697. 数组的度

给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。

你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。



示例 1：

输入：[1, 2, 2, 3, 1]
输出：2
解释：
输入数组的度是2，因为元素1和2的出现频数最大，均为2.
连续子数组里面拥有相同度的有如下所示:
[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
最短连续子数组[2, 2]的长度为2，所以返回2.

示例 2：

输入：[1,2,2,3,1,4,2]
输出：6



提示：

    nums.length 在1到 50,000 区间范围内。
    nums[i] 是一个在 0 到 49,999 范围内的整数。

*/

type FirstLastIndex struct {
	FirstIndex int //首次出现的下标
	LastIndex  int //最后出现的下标
	Count      int //出现的次数
}

/*
思路：遍历一次数组用map记录每个数下标及次数信息
遍历map找到出现最多次数的数字的最小长度
*/
func FindShortestSubArray(nums []int) int {
	result, max := len(nums), 0
	group := map[int]*FirstLastIndex{}

	for i, d := range nums {
		_, ok := group[d]
		if !ok {
			group[d] = &FirstLastIndex{
				FirstIndex: i,
				LastIndex:  i,
				Count:      1,
			}
		} else {
			group[d].LastIndex = i
			group[d].Count++
		}
	}

	for _, v := range group {
		if v.Count < max {
			continue
		}

		if v.Count > max {
			result = v.LastIndex - v.FirstIndex + 1
			max = v.Count
			continue
		}

		if v.LastIndex-v.FirstIndex+1 < result {
			result = v.LastIndex - v.FirstIndex + 1
		}

	}

	return result
}
