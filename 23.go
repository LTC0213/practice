package main

import (
	"container/heap"
	"fmt"
)

func mergeKLists(lists []*ListNode) *ListNode {
	mh := &minHeap{}
	for _, list := range lists {
		if list != nil {
			heap.Push(mh, list)
		}
	}
	dummy := &ListNode{Val: -1}
	p := dummy
	for mh.Len() > 0 {
		cur := heap.Pop(mh).(*ListNode)
		p.Next = cur
		p = p.Next
		if cur.Next != nil {
			heap.Push(mh, cur.Next)
		}
	}
	return dummy.Next
}

type minHeap []*ListNode

func (m minHeap) Len() int {
	return len(m)
}
func (m minHeap) Less(i, j int) bool {
	return m[i].Val < m[j].Val
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*ListNode))
}
func (m *minHeap) Pop() (x interface{}) {
	n := len(*m)
	x, *m = (*m)[n-1], (*m)[:n-1]
	return x
}

type ListNode struct {
	Val  int
	Next *ListNode
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
	var lists []*ListNode
	head := &ListNode{Val: 1}
	head.addNode(4, 5)
	lists = append(lists, head)
	head = &ListNode{Val: 1}
	head.addNode(3, 4)
	lists = append(lists, head)
	head = &ListNode{Val: 2}
	head.addNode(6)
	lists = append(lists, head)
	res := mergeKLists(lists)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
