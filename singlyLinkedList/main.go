package main

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
