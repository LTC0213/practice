package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		a, b := p.Next, p.Next.Next
		p.Next = b
		a.Next = b.Next
		b.Next = a
		p = a
	}
	return dummy.Next
}
