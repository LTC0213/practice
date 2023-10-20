package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	base := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			base += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			base += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: base % 10}
		p = p.Next
		base = base / 10
	}
	if base > 0 {
		p.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
