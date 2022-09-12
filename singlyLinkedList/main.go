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

// Singly Linked List Constructor
func NewSinglyLinkedList() *SinglyLinkedList {
}
