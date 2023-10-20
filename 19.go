package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cnt := 0
	p := head
	for p != nil {
		cnt++
		p = p.Next
	}
	cnt -= n
	dummy := &ListNode{Next: head}
	p = dummy
	for cnt > 0 {
		cnt--
		p = p.Next
	}
	p.Next = p.Next.Next
	return dummy.Next
}
