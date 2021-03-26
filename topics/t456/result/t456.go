package result

/*
456. 132 模式

给你一个整数数组 nums ，数组中共有 n 个整数。132 模式的子序列 由三个整数 nums[i]、nums[j] 和 nums[k] 组成，并同时满足：i < j < k 和 nums[i] < nums[k] < nums[j] 。

如果 nums 中存在 132 模式的子序列 ，返回 true ；否则，返回 false 。

进阶：很容易想到时间复杂度为 O(n^2) 的解决方案，你可以设计一个时间复杂度为 O(n logn) 或 O(n) 的解决方案吗？

示例 1：

输入：nums = [1,2,3,4]
输出：false
解释：序列中不存在 132 模式的子序列。

示例 2：

输入：nums = [3,1,4,2]
输出：true
解释：序列中有 1 个 132 模式的子序列： [1, 4, 2] 。

示例 3：

输入：nums = [-1,3,2,0]
输出：true
解释：序列中有 3 个 132 模式的的子序列：[-1, 3, 2]、[-1, 3, 0] 和 [-1, 2, 0] 。

提示：

    n == nums.length
    1 <= n <= 104
    -109 <= nums[i] <= 109


*/

// 穷举，时间不通过
func Find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	list := make([][]int, 0)

	for _, num := range nums {
		newList := make([][]int, 0)
		for _, pair := range list {
			if len(pair) == 2 {
				if num > pair[0] && num < pair[1] {
					return true
				}
			} else {
				if pair[0]+1 < num {
					newPair := []int{pair[0], num}
					newList = append(newList, newPair)
				}
			}
		}

		newPair := []int{num}
		newList = append(newList, newPair)
		list = append(list, newList...)
	}

	return false
}

// 思路:边遍历边记录可能的区间范围
func Find132pattern2(nums []int) bool {
	pairList := make([][]int, 0, 8)
	pair := []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		for _, p := range pairList {
			if p[0] < nums[i] && p[1] > nums[i] {
				return true
			}
		}

		if len(pair) == 1 {
			if pair[0]+1 < nums[i] {
				pair = append(pair, nums[i])
				pairList = append(pairList, pair)
				continue
			}

			if pair[0] >= nums[i] {
				pair[0] = nums[i]
			}
			continue
		}

		if len(pair) == 2 {
			if nums[i] >= pair[1] {
				pair[1] = nums[i]
				continue
			}

			if nums[i] < pair[0] {
				pair = []int{nums[i]}
			}
		}
	}

	return false
}
