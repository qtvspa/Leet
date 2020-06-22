package main

// 两数相加（逆序链表）
// https://leetcode-cn.com/problems/add-two-numbers/
//


type ListNode struct {
	Val int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headList := new(ListNode)
	head := headList
	num := 0

	for l1 != nil || l2 != nil || num > 0 {
		headList.Next = new(ListNode)
		headList = headList.Next

		if l1 != nil {
			num = num + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num = num + l2.Val
			l2 = l2.Next
		}
		// 取余数作为Val值
		headList.Val = (num) % 10

		num = num / 10
	}

	return head.Next
}