package main

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
