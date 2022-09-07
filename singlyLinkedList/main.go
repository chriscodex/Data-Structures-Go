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
