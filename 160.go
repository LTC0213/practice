package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}

func (node *ListNode) addNode(val int) {
	end := &ListNode{Val: val}
	for node.Next != nil {
		node = node.Next
	}
	node.Next = end
}
