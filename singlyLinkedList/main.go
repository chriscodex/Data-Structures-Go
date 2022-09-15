package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

/* Constructors */
// Singly Linked List Constructor
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Node constructor
func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

/*Methods*/
// Traverse the list
func (sll *SinglyLinkedList) Traversal() {
	pointer := sll.head
	for pointer != nil {
		fmt.Println(pointer.data)
		pointer = pointer.next
	}
}

// Add a node at the end of the list
func (sll *SinglyLinkedList) AddNodeEnd(data int) {
	node := NewNode(data)
	if sll.head == nil {
		sll.head = node
		sll.tail = node
	} else {
		sll.tail.next = node
		sll.tail = node
	}
}

/* Add a node at the top of the list */
func (sll *SinglyLinkedList) AddNodeBeggin(data int) {
	node := NewNode(data)
	if sll.head == nil {
		sll.head = node
		sll.tail = node
	} else {
		node.next = sll.head
		sll.head = node
	}
	sll.length++
}

func (sll *SinglyLinkedList) AddNodeSpecify(data int, position int) {
	node := NewNode(data)
	if position == 0 {
		sll.AddNodeBeggin(node.data)
	} else {
	}
	sll.length++
}
