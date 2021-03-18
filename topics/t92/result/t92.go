package result

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
92. 反转链表 II
给你单链表的头节点 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。

示例 1：

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

示例 2：

输入：head = [5], left = 1, right = 1
输出：[5]

提示：

    链表中节点数目为 n
    1 <= n <= 500
    -500 <= Node.val <= 500
    1 <= left <= right <= n

 */

// 思路：找到left节点与right节点，调换后找到下一个left节点和right节点，以此类推。注意left节点和right节点相邻情况要特殊处理
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if right == 1 || left >= right {
		return head
	}

	var leftPreNode, leftNode, rightPreNode, rightNode, nowNode, nowPreNode *ListNode

	nowNode, leftNode, rightNode = head, head, head
	for i := 2; i <= right; i++ {
		nowPreNode = nowNode
		nowNode = nowNode.Next

		if i == left {
			leftNode = nowNode
			leftPreNode = nowPreNode
		}

		if i == right {
			rightNode = nowNode
			rightPreNode = nowPreNode
		}
	}

	for i := 0; left+i < right-i; i++ {
		if leftPreNode == nil {
			head = rightNode
			leftNode.Next, rightNode.Next = rightNode.Next, leftNode.Next

			if rightNode == rightNode.Next {
				rightNode.Next = leftNode
			} else {
				rightPreNode.Next = leftNode
			}

			leftPreNode = head
			leftNode = head.Next
		} else {
			if leftNode.Next == rightNode{
				leftPreNode.Next = rightNode
				leftNode.Next = rightNode.Next
				rightNode.Next = leftNode
			} else {
				leftPreNode.Next, rightPreNode.Next = rightNode, leftNode
				leftNode.Next, rightNode.Next = rightNode.Next, leftNode.Next
			}

			leftPreNode, leftNode = rightNode, rightNode.Next
		}

		rightPreNode, rightNode = leftPreNode, leftNode
		for j := left + i + 1; j < right-i-1; j++ {
			rightPreNode, rightNode = rightNode, rightNode.Next
		}
	}

	return head
}
