package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	for k := 1; k < n; k *= 2 {
		cur := dummy
		for i := 0; i+k < n; i += 2 * k {
			left, right := cur.Next, cur.Next
			for j := 0; j < k && i+j < n; j++ {
				right = right.Next
			}
			l, r := 0, 0
			for l < k && r < k && right != nil {
				if left.Val <= right.Val {
					cur.Next = left
					cur = cur.Next
					left = left.Next
					l++
				} else {
					cur.Next = right
					cur = cur.Next
					right = right.Next
					r++
				}
			}
			for l < k {
				cur.Next = left
				cur = cur.Next
				left = left.Next
				l++
			}
			for r < k && right != nil {
				cur.Next = right
				cur = cur.Next
				right = right.Next
				r++
			}
			cur.Next = right
		}
	}
	return dummy.Next
}

func (head *ListNode) addNode(vals ...int) *ListNode {
	p := head
	for p.Next != nil {
		p = p.Next
	}
	for _, val := range vals {
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return head
}

func main() {
	head := &ListNode{Val: -1}
	head.addNode(5, 3, 4, 0)
	head = sortList(head)
	p := head
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}
