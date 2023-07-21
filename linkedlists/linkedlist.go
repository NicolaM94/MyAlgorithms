package linkedlists

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Add(value int) {
	if ll.Head == nil {
		ll.Head = &Node{Value: value}
		return
	}
	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Value: value}
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
	var prev *Node
	curr := ll.Head
	var next *Node
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	ll.Head = prev
}

func (ll *LinkedList) Next() Node {
	ll.Head = ll.Head.Next
	return *ll.Head
}

func (ll *LinkedList) Lenght() (lenght int) {
	current := ll.Head
	for current.Next != nil {
		lenght++
		current = current.Next
	}
	lenght++
	return
}
