package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	p := head
	for list1 != nil && list2 != nil {
		l1, l2 := list1.Val, list2.Val
		if l1 <= l2 {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next
}

func (node *ListNode) addNode(val int) *ListNode {
	head := node
	for node.Next != nil {
		node = node.Next
	}
	node.Next = &ListNode{Val: val}
	return head
}

func main() {
	list1 := &ListNode{Val: 1}
	list1.addNode(2).addNode(4)
	list2 := &ListNode{Val: 1}
	list2.addNode(3).addNode(4)
	res := mergeTwoLists(list1, list2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
