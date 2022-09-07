package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}

// Singly Linked List constructor
func NewSinglyLinkedList(data int) *SinglyLinkedList {
	n := Node{
		data: data,
	}
	return &SinglyLinkedList{
		head: &n,
		tail: &n,
	}
}

// Node constructor
func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

// Methods
func (sll *SinglyLinkedList) Traversal() {
	pointer := sll.head
	for pointer != nil {
		fmt.Println(pointer.data)
		pointer = pointer.next
	}
}

func (sll *SinglyLinkedList) AddNodeEnd(node *Node) {
	sll.tail.next = node
	sll.tail = node
}

func (sll *SinglyLinkedList) AddNodeBeggin(node *Node) {
	node.next = sll.head
	sll.head = node
}

func (sll *SinglyLinkedList) AddNodeSpecify(node *Node, position int) {
	if position != 0 {
		pointer := sll.head
		for i := 0; i < position-1; i++ {
			pointer = pointer.next
		}
		after := pointer.next
		node.next = after
		pointer.next = node
	} else {
		sll.AddNodeBeggin(node)
	}
}

func main() {
	slink := NewSinglyLinkedList(1)
	n2 := NewNode(2)
	n3 := NewNode(0)
	n4 := NewNode(3)
	n5 := NewNode(5)
	slink.AddNodeEnd(n2)
	slink.AddNodeBeggin(n3)
	slink.AddNodeSpecify(n4, 1)
	slink.AddNodeSpecify(n5, 0)
	slink.Traversal()
	for i := 0; i < 0; i++ {
		fmt.Println(i)
	}
}
