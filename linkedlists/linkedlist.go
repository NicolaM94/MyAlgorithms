package linkedlists

import "fmt"

type listNode struct {
	Value int
	Next  *listNode
}

type LinkedList struct {
	Head *listNode
}

func (ll *LinkedList) Add(value int) {
	if ll.Head == nil {
		ll.Head = &listNode{Value: value}
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &listNode{Value: value}
}

func (ll *LinkedList) Remove(target int) {
	if ll.Head.Value == target {
		ll.Head.Next = ll.Head.Next.Next
		return
	}
	prev := ll.Head
	current := ll.Head.Next
	for current.Next != nil {
		if current.Value == target {
			prev.Next = current.Next
			return
		}
		prev = current
		current = current.Next
	}
	if current.Value == target {
		prev.Next = nil
	}
}

func (ll *LinkedList) Print() {
	current := ll.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}

func (ll *LinkedList) Reverse() {
	var prev *listNode
	curr := ll.Head
	var next *listNode
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.Head = prev
}
