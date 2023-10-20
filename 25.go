package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		p = p.Next
	}
	var q *ListNode
	p = head
	for i := 0; i < k; i++ {
		tmp := p.Next
		p.Next = q
		q = p
		p = tmp
	}
	head.Next = reverseKGroup(p, k)
	return q
}
