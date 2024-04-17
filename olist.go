package main

import (
	"fmt"
)

type Crd struct {
	x, y int
}

type CrdNode struct {
	crd  Crd
	next *CrdNode
}

type CrdList struct {
	head *CrdNode
}

func (l *CrdList) Insert(x, y int) {
	new_crd := &CrdNode{Crd{x, y}, nil}

	if l.head == nil {
		l.head = new_crd
		return
	}

	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = new_crd
}

type Node struct {
	x, y, val  int
	prev, next *Node
}

func (n *Node) crd() Crd {
	return Crd{n.x, n.y}
}

type OList struct {
	head *Node
}

func (list *OList) Insert(x, y, val int, prev *Node) {
	new_node := &Node{x, y, val, prev, nil}

	if list.head == nil {
		list.head = new_node
	} else if list.head.val > val {
		new_node.next = list.head
		list.head = new_node
	} else {
		cur := list.head
		for {
			if cur.next == nil {
				cur.next = new_node
				break
			} else if cur.next.val > val {
				new_node.next = cur.next
				cur.next = new_node
				break
			}
			cur = cur.next
		}
	}
}

func (list *OList) Print() {
	cur := list.head
	for cur != nil {
		fmt.Printf("%d", cur.val)
		if cur.next != nil {
			fmt.Print("-")
		}
		cur = cur.next
	}
	fmt.Println()
}

func (list *OList) Pop() (*Node, error) {
	if list.head == nil {
		return nil, fmt.Errorf("no head")
	}
	r := list.head
	list.head = list.head.next
	return r, nil
}
