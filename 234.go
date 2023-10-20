package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	p1, p2 := head, reverse(slow.Next)
	slow.Next = nil
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func reverse(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}
	var prev *ListNode
	cur := node
	for cur != nil {
		cur.Next, cur, prev = prev, cur.Next, cur
	}
	return prev
}
