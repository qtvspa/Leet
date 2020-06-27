package main

// 移除一个链表中的重复结点（重复时保留第一个结点）
// https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
// 思路：用值的哈希表判断是否重复


type ListNode struct {
	Val int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 创建哈希表
	occurred := map[int]bool{head.Val: true}
	pos := head

	// 依次判定当前遍历下一个的结点是否重复
	for pos.Next != nil {
		cur := pos.Next
		if !occurred[cur.Val] {
			occurred[cur.Val] = true
			pos = pos.Next
		} else {
			// 重复则跳过
			pos.Next = pos.Next.Next
		}
	}
	pos.Next = nil
	return head
}


