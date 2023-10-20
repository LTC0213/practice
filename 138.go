package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	p := head
	for p != nil {
		tmp := p.Next
		p.Next = &Node{
			Val:  p.Val,
			Next: tmp,
		}
		p = p.Next.Next
	}

	p = head
	for p != nil {
		q := p.Next
		if p.Random != nil {
			q.Random = p.Random.Next
		}
		p = p.Next.Next
	}

	res := head.Next
	p = head
	for p != nil {
		q := p.Next
		p.Next = q.Next
		p = p.Next
		if p != nil {
			q.Next = p.Next
		}
	}
	return res
}
