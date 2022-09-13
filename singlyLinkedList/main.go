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
}
