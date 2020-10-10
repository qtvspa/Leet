package main

// 判定是否有环
// 思路: 哈希表、快慢指针

type ListNode struct {
	Val int
	Next *ListNode
}


func hasCycleHash(head *ListNode) bool {

	seen := map[*ListNode]int{}
	for head != nil {
		if _, ok := seen[head]; ok {
			return true
		}
		seen[head] = 1
		head = head.Next
	}
	return false
}


func hasCyclePoint(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow, fast = slow.Next, fast.Next.Next
	}
	return true
}
