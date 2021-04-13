package result

/*
783. 二叉搜索树节点最小距离

给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/ 相同



示例 1：

输入：root = [4,2,6,1,3]
输出：1

示例 2：

输入：root = [1,0,48,null,null,12,49]
输出：1

 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list []int = make([]int, 0)

func MinDiffInBST(root *TreeNode) int {
	list = make([]int, 0)
	InitList(root)

	result := list[1] - list[0]
	if result == 1 {
		return 1
	}

	for i := 2; i < len(list); i++ {
		diff := list[i] - list[i-1]

		if diff == 1 {
			return 1
		}

		if diff < result {
			result = diff
		}
	}
	return result
}

//左中右遍历树
func InitList(root *TreeNode) {
	if root.Left != nil {
		InitList(root.Left)
	}

	list = append(list, root.Val)

	if root.Right != nil {
		InitList(root.Right)
	}
}

